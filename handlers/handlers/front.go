package handlers

import (
	"fmt"
	"net/http"

	user "PhotoBlog/BL/user"
	ui "PhotoBlog/UI"
)

func Front(res http.ResponseWriter, req *http.Request) {
	ui.Front(res)
}

func Register(res http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	family := req.FormValue("family")
	email := req.FormValue("email")

	user := user.NewUser(name, family, email)

	if user.IsFalseUserDTO() {
		fmt.Println("could not create this user. email already exists?")

		return
	}

	fmt.Println("new user: ", user)

}
