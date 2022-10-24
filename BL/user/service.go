package user

import "github.com/google/uuid"

func NewUser(name string, family string) uuid.UUID {

	return uuid.New()
}
