package userapi

import (
	"backend-github-trending/errlog"
	"backend-github-trending/log"
	"backend-github-trending/model"
	"backend-github-trending/model/req"
	"backend-github-trending/repository"
	"backend-github-trending/security"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo"
)

// UserHandler for handling API like sign in, sign up
type UserHandler struct {
	UserRepo repository.UserRepo
}

// SignIn handler: POST("/user/signIn")
func (u UserHandler) SignIn(c echo.Context) error {
	req := req.SignInReq{}

	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user, err := u.UserRepo.CheckLogin(c.Request().Context(), req)
	if err != nil {
		return c.JSON(
			http.StatusUnauthorized,
			model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    errlog.ErrLogInFailed.Error(),
				Data:       nil,
			},
		)
	}

	isPasswordRight := security.ComparePasswords(user.Password, []byte(req.Password))
	if !isPasswordRight {
		return c.JSON(
			http.StatusUnauthorized,
			model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    errlog.ErrLogInFailed.Error(),
				Data:       nil,
			},
		)
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       user,
	})

}

// SignUp handler  POST("/user/signUp")
func (u UserHandler) SignUp(c echo.Context) error {
	req := req.SignUpReq{}

	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	hash := security.HashAndSalt([]byte(req.Password))
	role := model.MEMBER.ToString()

	userID, err := uuid.NewUUID()

	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user := model.User{
		UserID:   userID.String(),
		Email:    req.Email,
		FullName: req.FullName,
		Password: hash,
		Role:     role,
		Token:    "",
	}

	user, err = u.UserRepo.SaveUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       "",
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       user,
	})
}
