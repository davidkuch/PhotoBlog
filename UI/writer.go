package ui

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

//var viewToTemplate = map[string]string{
//	"front" : "index.html",
//	"user_home" : "user_home.html"
//}

func init() {
	tpl = template.Must(template.ParseGlob("./UI/templates/components/*.html"))
	template.Must(tpl.ParseGlob("./UI/templates/*.html"))

	fmt.Println("UI: templates parsed")
}

func ShowView(writer http.ResponseWriter, view string) {

	to_execute := view + ".html"

	fmt.Println("UI:to_execute: ", to_execute)

	tpl.ExecuteTemplate(writer, to_execute, nil)
}
