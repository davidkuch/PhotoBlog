package middleware

import (
	"fmt"
	"net/http"
)

func MakeSecure(raw func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		status := checkSession(req)

		if !status {
			//session err return
			fmt.Println("session err")
		}

		raw(res, req)

	}
}

func checkSession(req *http.Request) bool {
	cookie, err := req.Cookie("session-id")

	if err == http.ErrNoCookie {
		return false //temp: to be fixed later
	}

	return true
}
