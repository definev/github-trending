package repoimpl

import (
	"backend-github-trending/db"
	"backend-github-trending/errlog"
	"backend-github-trending/log"
	"backend-github-trending/model"
	"backend-github-trending/repository"
	"context"
	"time"

	"github.com/lib/pq"
)

// UserRepoImpl implement of UserRepo
type UserRepoImpl struct {
	sql *db.SQL
}

// SaveUser impl
func (u UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := `
		INSERT INTO users(user_id, email, password, role, full_name, created_at, updated_at)
		VALUES(:user_id, :email, :password, :role, :full_name, :created_at, :updated_at)
	`

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := u.sql.Db.NamedExecContext(context, statement, user)

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

// NewUserRepo create new user repo
func NewUserRepo(sql *db.SQL) repository.UserRepo {
	return &UserRepoImpl{sql: sql}
}
