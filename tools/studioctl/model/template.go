package model

import "fmt"

var tpl_model = fmt.Sprintf(`package model
{{if .HasImport}}
import "go-studio/core/jsontime"
{{end}}
type (
	{{.TableMapperName}}Model struct {
		db *xorm.Engine
	}
	
	{{.TableMapperName}} struct {
{{range .Columns}}		{{.MapperName}}	{{.Type}}	%s{{.Tag}}%s 
{{end}}	}
)

func New{{.TableMapperName}}Model(db *xorm.Engine) *{{.TableMapperName}}Model {
	return &{{.TableMapperName}}Model{db:db}
}

func ({{.TableMapperName}}) TableName() string {
	return "{{.TableName}}"
}

func (m *{{.TableMapperName}}Model) Insert(data *{{.TableMapperName}}) error {
	affect, err := m.db.InsertOne(data)
	if err != nil {
		return err
	}
	if affect != 1 {
		return errorx.DBUpdateNotAffected
	}
	return nil
}

func (m *{{.TableMapperName}}Model) FindOne(id int64) (*{{.TableMapperName}}, error) {
	var resp {{.TableMapperName}}
	has, err := m.db.ID(id).Get(&resp)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errorx.DBDataNotFound
	}
	return &resp, nil
}

func (m *{{.TableMapperName}}Model) Update(data *{{.TableMapperName}}) error {
	affect, err := m.db.ID(data.Id).Update(data)
	if err != nil {
		return err
	}

	if affect != 1 {
		return errorx.DBUpdateNotAffected
	}
	return nil
}

func (m *{{.TableMapperName}}Model) Delete(id int64) error {
	affect, err := m.db.ID(id).Delete({{.TableMapperName}}{})
	if err != nil {
		return err
	}
	if affect != 1 {
		return errorx.DBUpdateNotAffected
	}
	return nil
}



`, "`", "`")
