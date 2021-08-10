package model

import (
	"bytes"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/urfave/cli"
	"html/template"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
	"xorm.io/xorm/schemas"
)

const (
	flagURL   = "url"
	flagDir   = "dir"
	flagTable = "table"
	flagStyle = "style"
)

type TableMapper struct {
	TableName       string
	TableMapperName string
	Columns         []ColumnMapper
	HasImport       bool
}
type ColumnMapper struct {
	Name       string
	MapperName string
	Type       string
	Comment    string
	Tag        template.HTML
}

func MysqlDataSource(ctx *cli.Context) error {
	url := ctx.String(flagURL)
	dir := ctx.String(flagDir)
	table := ctx.String(flagTable)
	style := ctx.String(flagStyle)
	return fromDataSource(url, table, dir, style)
}

func fromDataSource(url, tableName, dir, style string) error {
	if dir != "" {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Println("mkdir err:" + err.Error())
				return err
			}
		}
	}

	current, err := user.Current()
	if err != nil {
		return err
	}
	tpl := tpl_model
	var mapper names.Mapper

	if style == "same" {
		mapper = names.SameMapper{}
	} else {
		mapper = names.SnakeMapper{}
	}

	if isExist(current.HomeDir + "/.studioctl/model.template") {
		file, err := os.Open(current.HomeDir + "/.studioctl/model.template")
		if err != nil {
			return err
		}
		all, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		tpl = fmt.Sprintf(string(all), "`", "`")
	}

	engine, err := xorm.NewEngine("mysql", url)
	if err != nil {
		return err
	}

	t := template.New("gen model")
	tmpl, err := t.Parse(tpl)
	if err != nil {
		return err
	}

	dbMetas, err := engine.DBMetas()
	if err != nil {
		fmt.Println("DBMetas:" + err.Error())
		return err
	}

	for _, table := range dbMetas {
		match, err := filepath.Match(tableName, table.Name)
		if err != nil {
			return err
		}
		if match {
			tableMapper := TableMapper{
				TableName:       table.Name,
				TableMapperName: names.SnakeMapper{}.Table2Obj(table.Name),
				Columns:         make([]ColumnMapper, 0, len(table.ColumnsSeq())),
			}

			if isExist(path.Join(dir, fmt.Sprintf("%s.go", tableMapper.TableName))) {
				fmt.Println(table.Name + " is exist")
				continue
			}

			for _, column := range table.Columns() {
				if typeString(column) == "jsontime.JsonTime" {
					tableMapper.HasImport = true
				}
				tableMapper.Columns = append(tableMapper.Columns, ColumnMapper{
					Name:       column.Name,
					MapperName: mapper.Table2Obj(column.Name),
					Type:       typeString(column),
					Comment:    column.Comment,
					Tag:        tag(table, column),
				})
			}
			newbytes := bytes.NewBufferString("")
			err = tmpl.Execute(newbytes, tableMapper)
			all, _ := ioutil.ReadAll(newbytes)

			file, err := os.Create(path.Join(dir, tableMapper.TableName+".go"))
			if err != nil {
				return err
			}
			file.Write(all)
			file.Close()
		}
	}

	return nil
}

func typeString(col *schemas.Column) string {
	switch col.SQLType.Name {
	case schemas.UnsignedBigInt, schemas.UnsignedInt:
		return reflect.TypeOf(int64(1)).String()
	case schemas.UnsignedTinyInt, schemas.UnsignedMediumInt, schemas.UnsignedSmallInt:
		return reflect.TypeOf(1).String()
	case schemas.Decimal:
		return reflect.TypeOf(float64(1)).String()
	}

	s := schemas.SQLType2Type(col.SQLType).String()
	if s == "time.Time" {
		return "jsontime.JsonTime"
	}
	return s
}

func tag(table *schemas.Table, col *schemas.Column) template.HTML {

	var res []string

	if col.IsPrimaryKey {
		res = append(res, "pk")
	}

	if col.IsAutoIncrement {
		res = append(res, "autoincr")
	}

	if col.SQLType.IsTime() && strings.Contains(col.Name, "created") {
		res = append(res, "created")
	}

	if col.SQLType.IsTime() && strings.Contains(col.Name, "updated") {
		res = append(res, "updated")
	}

	if col.SQLType.IsTime() && strings.Contains(col.Name, "delete") {
		res = append(res, "deleted")
	}

	res = append(res, "'"+col.Name+"'")

	names := make([]string, 0, len(col.Indexes))
	for name := range col.Indexes {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		index := table.Indexes[name]
		var uistr string
		if index.Type == schemas.UniqueType {
			uistr = "unique"
		} else if index.Type == schemas.IndexType {
			uistr = "index"
		}
		if len(index.Cols) > 1 {
			uistr += "(" + index.Name + ")"
		}
		res = append(res, uistr)
	}

	if len(res) > 0 {
		return template.HTML(fmt.Sprintf(`xorm:"%s"`, strings.Join(res, " ")))
	}
	return ""
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
