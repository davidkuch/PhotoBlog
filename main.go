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

	mux := http.NewServeMux()

	mux.Handle("/front", http.HandlerFunc(front))

	http.ListenAndServe(":3000", mux)
}
