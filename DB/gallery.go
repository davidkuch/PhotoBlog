package db

import (
	"PhotoBlog/model"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SaveNewGallery(gal *model.GalleryDTO) {
	db, err := sql.Open("mysql", open_stt)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := fmt.Sprintf("INSERT INTO galleries(id,name,created,last_edit) VALUES (%s, %s, %s, %s, %s)", gal.Uid, gal.Name, gal.Created, gal.Last_edit)

	db.Exec(sql)
}
