package user

import (
	"PhotoBlog/model"

	"github.com/google/uuid"
)

type user struct {
	id     uuid.UUID
	name   string
	family string
}

func wrap(src user) model.User {
	return model.User{src.id, src.name, src.family}
}

func unwrap(src model.User) user {
	return user{src.Id, src.Name, src.Family}
}
