package model

import (
	"time"

	"github.com/google/uuid"
)

type GalleryDTO struct {
	Uid       uuid.UUID
	Name      string
	Created   time.Time
	Last_edit time.Time
	Pics      []uuid.UUID
}
