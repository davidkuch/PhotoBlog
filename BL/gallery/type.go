package gallery

import (
	"PhotoBlog/model"
	"time"

	"github.com/google/uuid"
)

type gallery struct {
	uid       uuid.UUID
	name      string
	created   time.Time
	last_edit time.Time
	pics      []uuid.UUID
}

func unwrap(dto *model.GalleryDTO) *gallery {
	return &gallery{dto.Uid, dto.Name, dto.Created, dto.Last_edit, dto.Pics}

}

func wrap(data *gallery) *model.GalleryDTO {
	return &model.GalleryDTO{data.uid, data.name, data.created, data.last_edit, data.pics}
}
