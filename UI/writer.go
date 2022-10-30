package ui

import (
	"fmt"
	"html/template"
	"net/http"
)

var Tpl *template.Template

func init() {
	Tpl = template.Must(template.ParseGlob("./UI/templates/*.html"))
	fmt.Println("UI: templates parsed")
}

func Front(writer http.ResponseWriter) {

	Tpl.ExecuteTemplate(writer, "index.html", nil)
}
