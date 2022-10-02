package handlers

import (
	"PhotoBlog/BL/gallery"
	ui "PhotoBlog/UI"
	"net/http"
)

func Front(res http.ResponseWriter, req *http.Request) {
	gal := gallery.GetGallery()

	ui.ShowGallery(res, gal)

}
