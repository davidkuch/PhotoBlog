package ui

import (
	"fmt"
	"html/template"
)

var Tpl *template.Template

func init() {
	Tpl = template.Must(template.ParseGlob("./UI/templates/*.html"))
	fmt.Println("UI: templates parsed")
}
