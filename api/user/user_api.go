package userapi

import (
	routermodel "backend-github-trending/model/router"
)

// UserAPI collection of user service
type UserAPI struct {
	UserHandler UserHandler
	UserRouter  UserRouter
}

// Init UserAPI
func (api *UserAPI) Init() {
	api.UserRouter.signIn = routermodel.Router{
		Name:     "/signIn",
		Method:   routermodel.POST,
		Function: api.UserHandler.SignIn,
	}
	api.UserRouter.signIn.SetupRouter()

	api.UserRouter.signUp = routermodel.Router{
		Name:     "/signUp",
		Method:   routermodel.POST,
		Function: api.UserHandler.SignUp,
	}
	api.UserRouter.signUp.SetupRouter()
}
