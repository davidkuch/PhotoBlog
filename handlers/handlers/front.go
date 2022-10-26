package handlers

import (
	"PhotoBlog/BL/gallery"
	ui "PhotoBlog/UI"
	"net/http"

	"github.com/google/uuid"
)

func Front(res http.ResponseWriter, req *http.Request) {
	gal := gallery.GetGallery(uuid.New())

	ui.ShowGallery(res, gal)

}
