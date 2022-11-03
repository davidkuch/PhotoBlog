package main

import (
	"PhotoBlog/handlers"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("handlers: ", handlers.HandlersMapping)

	//register handlers for routs.
	for rout, handler := range handlers.HandlersMapping {
		http.Handle(rout, handler)
	}

	fmt.Println("Hello PhotoBlog!")

	http.ListenAndServe(":3000", nil)
}
