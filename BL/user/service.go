package user

import (
	"PhotoBlog/model"
	"fmt"

	"github.com/google/uuid"

	db "PhotoBlog/DB"
)

func NewUser(name, family, password, email string) model.UserDTO {
	new_id := uuid.New()

	new_user := user{new_id, name, family, email}

	//here goes some logic... validation etc'
	if userExists(new_user) {
		fmt.Println("user already exists with email: ", new_user.email)

		return model.MakeFalseUserDTO()
	}

	//and back to DTO and to other layer

	dto := wrap(new_user)

	db.SaveNewUser(dto, password)

	return dto
}
