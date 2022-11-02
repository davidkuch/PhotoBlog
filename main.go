package main

import (
	"PhotoBlog/handlers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello PhotoBlog!")

	mux := http.NewServeMux()

	fmt.Println(handlers.HandlersMapping)

	//register handlers for routs.
	for rout, handler := range handlers.HandlersMapping {
		mux.Handle(rout, handler)
	}

	http.ListenAndServe(":3000", mux)
}
