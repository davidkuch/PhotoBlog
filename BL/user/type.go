package user

import (
	"PhotoBlog/model"

	"github.com/google/uuid"
)

type user struct {
	id     uuid.UUID
	name   string
	family string
	email  string
}

func wrap(src user) model.UserDTO {
	return model.UserDTO{src.id, src.name, src.family, src.email}
}

func unwrap(src model.UserDTO) user {
	return user{src.Id, src.Name, src.Family, src.Email}
}
