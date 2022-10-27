package user

import (
	"PhotoBlog/model"

	"github.com/google/uuid"

	db "PhotoBlog/DB"
)

func NewUser(name string, family string) model.User {
	new_id := uuid.New()

	new_user := user{new_id, name, family}

	ret := wrap(new_user)

	db.SaveNewUser(ret)

	return ret
}
