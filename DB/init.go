package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Dummy() {
	fmt.Println("dummy!")
}

func init() {
	init_open_stt()
	ping()
	init_db()
}

func init_open_stt() string {
	config, err := ioutil.ReadFile("/home/davidk/github/PhotoBlog/DB/config.json")

	if err != nil {
		panic(err)
	}

	var config_map map[string]string

	json.Unmarshal([]byte(config), &config_map)

	username := config_map["username"]
	password := config_map["password"]

	open_stt := fmt.Sprintf("%s:%s@/photo_db?multiStatements=true", username, password)

	return open_stt
}

func ping() {

	fmt.Println("DB: attempt to ping...")

	open_stt := init_open_stt()

	var err error

	db, err = sql.Open("mysql", open_stt)

	if err != nil {
		panic(err)
	}
	fmt.Println("#2 db pointer to: ", db)

	db.Ping()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("DB:ping completed!")
}

func init_db() {

	sql, err := ioutil.ReadFile("/home/davidk/github/PhotoBlog/DB/init_data.sql")

	if err != nil {
		panic(err)
	}

	fmt.Println("#3 db pointer to: ", db)

	res, err := db.Exec(string(sql))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("DB: init completed!", res)

}
