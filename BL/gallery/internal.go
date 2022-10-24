package gallery

import "github.com/google/uuid"

func (*Gallery) save() uuid.UUID {

	return uuid.New()
}
