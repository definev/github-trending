package middleware

import (
	"backend-github-trending/model"
	"net/http"

	"github.com/labstack/echo"
)

// IsLogin middleware for check user is already login yet
func IsLogin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if 1 == 1 {
				return c.JSON(
					http.StatusPreconditionFailed,
					model.Response{
						StatusCode: http.StatusPreconditionFailed,
						Message:    "Failed",
						Data:       nil,
					},
				)
			}
			return next(c)
		}
	}

}
