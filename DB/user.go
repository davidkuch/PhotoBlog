package db

import (
	"PhotoBlog/model"
	"fmt"
)

func SaveNewUser(to_save model.User) {

	sql := "INSERT INTO users (id, name,family) VALUES (?,?,?)"

	result, err := db.Exec(sql, to_save.Id, to_save.Name, to_save.Family)

	if err != nil {
		fmt.Println("result: ", result)
		panic(err)

	}

}
