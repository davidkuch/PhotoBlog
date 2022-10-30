package main

import (
	"PhotoBlog/handlers/handlers"
	"PhotoBlog/handlers/middleware"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello PhotoBlog!")

	front := middleware.MakeSecure(handlers.Front)
	register := middleware.MakeSecure(handlers.Register)

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(front))
	mux.Handle("/register", http.HandlerFunc(register))

	http.ListenAndServe(":3000", mux)
}
