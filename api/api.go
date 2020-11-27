package api

import (
	userapi "backend-github-trending/api/user"
	routermodel "backend-github-trending/model/router"
	"fmt"

	"github.com/labstack/echo"
)

type baseAPI interface {
	Init()
}

// API : Instance of API
type API struct {
	Echo *echo.Echo
	User userapi.UserAPI
}

// Init api
func (api *API) Init() {
	routermodel.E = api.Echo
	api.User.Init()

	fmt.Print(api.Echo.Start(":8000"))
}

// InitAPI setup router and api
func InitAPI() {

}
