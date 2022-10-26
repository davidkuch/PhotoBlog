package ui

import (
	"PhotoBlog/model"
	"fmt"
	"html/template"
	"net/http"
)

var Tpl *template.Template

func init() {
	Tpl = template.Must(template.ParseGlob("./UI/templates/*.html"))
	fmt.Println("UI: templates parsed")
}

func ShowGallery(writer http.ResponseWriter, to_show *model.GalleryDTO) {
	Tpl.ExecuteTemplate(writer, "index.html", *to_show)
}
