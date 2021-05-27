package model

import "fmt"

var tpl_model = fmt.Sprintf(`package model
{{if .HasImport}}
import "go-studio/core/jsontime"
{{end}}
type {{.TableName}} struct {
{{range .Columns}}	{{.Name}}	{{.Type}}	%s{{.Tag}}%s 
{{end}}}



`, "`", "`")
