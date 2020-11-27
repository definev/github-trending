package userapi

import routermodel "backend-github-trending/model/router"

// UserRouter router
type UserRouter struct {
	signIn routermodel.Router
	signUp routermodel.Router
}
