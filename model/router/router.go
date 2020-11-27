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
	Name     string
	Method   string
	Function func(echo.Context) error
}

// SetupRouter for init Name in main.go
func (rest *Router) SetupRouter() {
	var Name *echo.Route

	switch {
	case rest.Method == GET:
		Name = E.GET(rest.Name, rest.Function)
	case rest.Method == POST:
		Name = E.POST(rest.Name, rest.Function)
	}

	if Name == nil {
		log.Errorf("%v have no method", rest.Name)
	}
}
