package gallery

import (
	db "PhotoBlog/DB"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func NewGallery(name string) uuid.UUID {
	gallery_uid := uuid.New()

	if name == "" {
		name = "unnamed gallery"
	}

	gallery := Gallery{gallery_uid, name, time.Now(), time.Now(), nil}

	fmt.Println("created gallery: ", gallery)

	return gallery_uid
}

func GetGallery(id uuid.UUID) *Gallery {
	fmt.Println("gallery service called")

	db.Dummy()

	return &Gallery{uuid.New(), "some", time.Now(), time.Now(), nil}
}
