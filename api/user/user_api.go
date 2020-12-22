package userapi

import (
	gtmiddleware "backend-github-trending/middleware"
	routermodel "backend-github-trending/model/router"

	"github.com/labstack/echo"
)

// UserAPI collection of user service
type UserAPI struct {
	UserHandler UserHandler
	UserRouter  UserRouter
}

// Init UserAPI
func (api *UserAPI) Init() {
	api.UserRouter.signIn = routermodel.Router{
		Name:       "/user/signIn",
		Method:     routermodel.POST,
		Function:   api.UserHandler.SignIn,
		Middleware: []echo.MiddlewareFunc{gtmiddleware.IsLogin()},
	}
	api.UserRouter.signIn.SetupRouter()

	api.UserRouter.signUp = routermodel.Router{
		Name:     "/user/signUp",
		Method:   routermodel.POST,
		Function: api.UserHandler.SignUp,
	}
	api.UserRouter.signUp.SetupRouter()
}
