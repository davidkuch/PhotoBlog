package middleware

import (
	"fmt"
	"net/http"
)

func Secure(raw func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("secure before")
		raw(res, req)
		fmt.Println("secure after")
	}
}
