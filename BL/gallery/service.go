package gallery

import (
	db "PhotoBlog/DB"
	"PhotoBlog/model"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func NewGallery(name string) uuid.UUID {
	gallery_uid := uuid.New()

	if name == "" {
		name = "unnamed gallery"
	}

	gallery := gallery{gallery_uid, name, time.Now(), time.Now(), nil}

	db.SaveNewGallery(wrap(&gallery))

	fmt.Println("created gallery: ", gallery)

	return gallery_uid
}

func GetGallery(id uuid.UUID) *model.GalleryDTO {
	fmt.Println("gallery service called")

	db.Dummy()

	return wrap(&gallery{uuid.New(), "some", time.Now(), time.Now(), nil})
}
