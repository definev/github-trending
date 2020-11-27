package main

import (
	"backend-github-trending/api"
	userapi "backend-github-trending/api/user"
	"backend-github-trending/db"
	"backend-github-trending/log"
	repoimpl "backend-github-trending/repository/repo_impl"
	"os"

	"github.com/labstack/echo"
)

func main() {
	sql := db.SQL{
		Host:     "localhost",
		Port:     3000,
		DbName:   "golang",
		Username: "postgres",
		Password: "admin",
	}

	sql.Connect()
	defer sql.Close()

	os.Setenv("APP_NAME", "github_trending")
	log.InitLogger(false)

	API := api.API{
		Echo: echo.New(),
		User: userapi.UserAPI{
			UserHandler: userapi.UserHandler{UserRepo: repoimpl.NewUserRepo(&sql)},
			UserRouter:  userapi.UserRouter{},
		},
	}
	API.Init()
}
