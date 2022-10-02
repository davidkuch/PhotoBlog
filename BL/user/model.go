package user

import "github.com/google/uuid"

type User struct {
	uid    uuid.UUID
	name   string
	family string
}
