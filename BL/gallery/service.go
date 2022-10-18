package gallery

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func GetGallery() *Gallery {
	fmt.Println("gallery service called")

	return &Gallery{uuid.New(), uuid.New(), "some", time.Now(), time.Now(), nil}
}

func NewGallery(name string) uuid.UUID {
	gallery_uid := uuid.New()

	if name == "" {
		name = "unnamed gallery"
	}

	gallery := Gallery{gallery_uid, uuid.Nil, name, time.Now(), time.Now(), nil}

	fmt.Println("creater gallery: ", gallery)

	return gallery_uid
}
