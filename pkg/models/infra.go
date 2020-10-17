package models

import "net/http"

//Route Type
type Route struct {
	MethodName      string
	Path            string
	HandlerFunction func(http.ResponseWriter, *http.Request)
	IsPublic        bool
}
