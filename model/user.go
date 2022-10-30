package model

import "github.com/google/uuid"

type UserDTO struct {
	Id     uuid.UUID
	Name   string
	Family string
	Email  string
}

func MakeFalseUserDTO() UserDTO {
	return UserDTO{uuid.Nil, "no-name", "no-family", "no-email"}
}

func (to_check *UserDTO) IsFalseUserDTO() bool {
	return *to_check == UserDTO{uuid.Nil, "no-name", "no-family", "no-email"}
}
