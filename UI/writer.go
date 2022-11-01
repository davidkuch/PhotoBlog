package ui

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./UI/templates/components/*.html"))
	template.Must(tpl.ParseGlob("./UI/templates/*.html"))

	fmt.Println("UI: templates parsed")
}

func Front(writer http.ResponseWriter) {

	tpl.ExecuteTemplate(writer, "index.html", nil)
}
