package middleware

import (
	"fmt"
	"net/http"
)

func MakeSecure(endpoint func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		status := checkSession(req)

		if !status {
			//session err return
			fmt.Println("session err")
		}
		//fmt.Println("handler called...")

		endpoint(res, req)

	}
}

func checkSession(req *http.Request) bool {
	cookie, err := req.Cookie("session-id")

	if err == http.ErrNoCookie {
		return true //temp: to be fixed later
	}

	if cookie.String() == "-1" {
		return false
	}

	return true
}
