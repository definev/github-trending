package repository

import (
	"backend-github-trending/model"
	"backend-github-trending/model/req"
	"context"
)

// UserRepo interface
type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	CheckLogin(context context.Context, req req.SignInReq) (model.User, error)
}
