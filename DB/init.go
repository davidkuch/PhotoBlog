package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var open_stt string

func Dummy() {
	fmt.Println("dummy!")
}

func init() {
	init_open_stt()
	ping()
	init_db()
}

func init_open_stt() {
	config, err := ioutil.ReadFile("/home/davidk/github/PhotoBlog/DB/config.json")

	if err != nil {
		panic(err)
	}

	var config_map map[string]string

	json.Unmarshal([]byte(config), &config_map)

	username := config_map["username"]
	password := config_map["password"]

	open_stt = fmt.Sprintf("%s:%s@/photo_db?multiStatements=true", username, password)
}

func ping() {

	fmt.Println("DB: attempt to ping")

	db, err := sql.Open("mysql", open_stt)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("DB:ping completed!")
}

func init_db() {

	db, err := sql.Open("mysql", open_stt)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	sql, err := ioutil.ReadFile("/home/davidk/github/PhotoBlog/DB/init_data.sql")

	if err != nil {
		panic(err)
	}

	res, err := db.Exec(string(sql))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("DB: init completed!", res)

}
