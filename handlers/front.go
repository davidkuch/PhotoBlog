package handlers

import (
	"fmt"
	"net/http"

	user "PhotoBlog/BL/user"
	ui "PhotoBlog/UI"
)

func front(res http.ResponseWriter, req *http.Request) {
	fmt.Println("front called!")
	ui.ShowView(res, "index")
}

func register(res http.ResponseWriter, req *http.Request) {
	fmt.Println("register called!")

	name := req.FormValue("name")
	family := req.FormValue("family")
	password := req.FormValue("password")
	email := req.FormValue("email")

	user := user.NewUser(name, family, password, email)

	if user.IsFalseUserDTO() {
		fmt.Println("could not create this user. email already exists?")

		return
	}

	fmt.Println("new user: ", user)

	res.WriteHeader(http.StatusCreated)

}
