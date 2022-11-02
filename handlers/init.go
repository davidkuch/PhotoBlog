package handlers

import (
	"PhotoBlog/handlers/middleware"
	"fmt"
	"net/http"
)

var HandlersMapping = map[string]http.Handler{}

// mapping of routs to raw handlers, before any decoration
// yet unused. waiting for refactor
var rawHandlersMapping = map[string]func(http.ResponseWriter, *http.Request){

	"/":        front,
	"register": register,
}

// decorates the handlers in handlersMapping
func init() {

	//refactor to use range over rawHandlersMapping.
	// *** beware of the for loop reference gotcha!... ***
	HandlersMapping["/"] = http.HandlerFunc(middleware.MakeSecure(front))
	HandlersMapping["/register"] = http.HandlerFunc(middleware.MakeSecure(register))

	fmt.Println(HandlersMapping)
}
