package routermodel

import (
	"backend-github-trending/log"

	"github.com/labstack/echo"
)

// E instance copy from router package
var E *echo.Echo

// GET method
var GET string = "GET"

// POST method
var POST string = "POST"

// Router struct
type Router struct {
	Name       string
	Method     string
	Function   func(echo.Context) error
	Middleware []echo.MiddlewareFunc
}

// SetupRouter for init Name in main.go
func (rest *Router) SetupRouter() {
	success := false

	switch {
	case rest.Method == GET:
		success = true
		if rest.Middleware != nil {
			E.GET(rest.Name, rest.Function, rest.Middleware...)
		} else {
			E.GET(rest.Name, rest.Function)
		}
	case rest.Method == POST:
		success = true
		if rest.Middleware != nil {
			E.POST(rest.Name, rest.Function, rest.Middleware...)
		} else {
			E.POST(rest.Name, rest.Function)
		}
	}

	if success == false {
		log.Errorf("%v have no method", rest.Name)
	}
}
