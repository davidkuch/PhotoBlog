package handlers

import (
	"fmt"
	"net/http"

	"PhotoBlog/BL/user"
)

func Front(res http.ResponseWriter, req *http.Request) {
	user := user.NewUser("diff", "folky")

	fmt.Println("user: ", user)

}
