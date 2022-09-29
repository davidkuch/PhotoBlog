package handlers

import (
	ui "PhotoBlog/UI"
	"net/http"
)

func Front(res http.ResponseWriter, req *http.Request) {
	ui.Tpl.ExecuteTemplate(res, "index.html", nil)

}
