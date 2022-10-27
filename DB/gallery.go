package db

import (
	"PhotoBlog/model"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SaveNewGallery(gal *model.GalleryDTO) {

	sql := fmt.Sprintf("INSERT INTO galleries(id,name,created,last_edit) VALUES (%s, %s, %s, %s, %s)", gal.Uid, gal.Name, gal.Created, gal.Last_edit)

	result, err := db.Exec(sql)

	fmt.Println("result, err = ", result, err)

}
