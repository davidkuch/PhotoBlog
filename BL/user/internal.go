package user

import (
	db "PhotoBlog/DB"
)

func userExists(user user) bool {
	email := user.email

	exists := db.EmailExists(email)

	return exists
}
