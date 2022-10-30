package db

import (
	"database/sql"

	"PhotoBlog/model"
	"fmt"
)

func SaveNewUser(to_save model.UserDTO) {

	stt := "INSERT INTO users (id, name,family,email) VALUES (?,?,?,?)"

	result, err := db.Exec(stt, to_save.Id, to_save.Name, to_save.Family, to_save.Email)

	if err != nil {
		fmt.Println("result: ", result)
		panic(err)

	}

}

func EmailExists(email string) bool {
	stt := "SELECT * FROM users WHERE email=?"

	result := db.QueryRow(stt, email)

	check := result.Scan()

	if check == sql.ErrNoRows {
		return false
	}

	return true

}
