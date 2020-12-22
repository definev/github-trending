package repoimpl

import (
	"backend-github-trending/db"
	"backend-github-trending/errlog"
	"backend-github-trending/log"
	"backend-github-trending/model"
	"backend-github-trending/model/req"
	"backend-github-trending/repository"
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

// UserRepoImpl implement of UserRepo
type UserRepoImpl struct {
	sql *db.SQL
}

// SaveUser insert User into postgres
func (u UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := `
		INSERT INTO users(user_id, email, password, role, full_name, created_at, updated_at)
		VALUES(:user_id, :email, :password, :role, :full_name, :created_at, :updated_at)
	`

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := u.sql.DB.NamedExecContext(context, statement, user)

	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, errlog.ErrUserExist
			}
		}
		return user, errlog.ErrSignUpFailed
	}

	return user, nil
}

// CheckLogin for checking login state
func (u *UserRepoImpl) CheckLogin(context context.Context, req req.SignInReq) (model.User, error) {
	var user model.User = model.User{}

	err := u.sql.DB.GetContext(context, &user, "SELECT * FROM users WHERE email=$1", req.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errlog.ErrUserNotFound
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}

// NewUserRepo create new user repo
func NewUserRepo(sql *db.SQL) repository.UserRepo {
	return &UserRepoImpl{sql: sql}
}
