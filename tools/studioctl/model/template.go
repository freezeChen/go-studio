package model

import "fmt"

var tpl_model = fmt.Sprintf(`package model

import "xorm.io/xorm"
{{if .HasImport}}
import "github.com/freezeChen/go-studio/core/jsontime"
{{end}}
type (
	{{.TableMapperName}}Model struct {
		db *xorm.Engine
		session *xorm.Session
	}
	
	{{.TableMapperName}} struct {
{{range .Columns}}		{{.MapperName}}	{{.Type}}	%s{{.Tag}} json:"{{.Name}}"%s // {{.Comment}}
{{end}}	}
)

func New{{.TableMapperName}}Model(db *xorm.Engine) *{{.TableMapperName}}Model {
	return &{{.TableMapperName}}Model{db:db}
}

func ({{.TableMapperName}}) TableName() string {
	return "{{.TableName}}"
}

func (m *{{.TableMapperName}}Model) WithSession(session *xorm.Session) *{{.TableMapperName}}Model {
	return &{{.TableMapperName}}Model{db:m.db, session: session}
}


func (m *{{.TableMapperName}}Model) Insert(data *{{.TableMapperName}}) error {
	var session =m.session
	if session ==nil {
		session = m.db.NewSession()
		defer session.Close()
	}

	affect, err := session.InsertOne(data)
	if err != nil {
		return err
	}
	if affect != 1 {
		return errcode.DBUpdateNotAffected
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
		return nil, errcode.DBDataNotFound
	}
	return &resp, nil
}

func (m *{{.TableMapperName}}Model) Update(data *{{.TableMapperName}}) error {
	var session =m.session
	if session ==nil {
		session = m.db.NewSession()
		defer session.Close()
	}

	_, err := session.ID(data.Id).Update(data)
	if err != nil {
		return err
	}

	return nil
}

func (m *{{.TableMapperName}}Model) Delete(id int64) error {
	var session =m.session
	if session ==nil {
		session = m.db.NewSession()
		defer session.Close()
	}

	affect, err := session.ID(id).Delete({{.TableMapperName}}{})
	if err != nil {
		return err
	}
	if affect != 1 {
		return errcode.DBUpdateNotAffected
	}
	return nil
}

func (m *{{.TableMapperName}}Model) List() (list []*{{.TableMapperName}}, err error) {
	list = make([]*{{.TableMapperName}}, 0)
	err = m.db.Find(&list)
	return
}
`, "`", "`")
