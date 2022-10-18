package gallery

import (
	"time"

	"github.com/google/uuid"
)

type Gallery struct {
	uid       uuid.UUID
	owner     uuid.UUID
	name      string
	created   time.Time
	last_edit time.Time
	pics      []uuid.UUID
}
