package ui

import (
	"fmt"
	"html/template"
	"net/http"

	"PhotoBlog/BL/gallery"
)

var Tpl *template.Template

func init() {
	Tpl = template.Must(template.ParseGlob("./UI/templates/*.html"))
	fmt.Println("UI: templates parsed")
}

func ShowGallery(writer http.ResponseWriter, to_show *gallery.Gallery) {
	Tpl.ExecuteTemplate(writer, "index.html", *to_show)
}
