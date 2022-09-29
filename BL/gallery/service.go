package gallery

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func GetGallery() *Gallery {
	fmt.Println("gallery service called")

	return &Gallery{uuid.New(), "some", time.Now(), time.Now(), nil}
}
