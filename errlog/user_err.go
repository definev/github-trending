package errlog

import "errors"

var (
	// ErrUserExist err code
	ErrUserExist = errors.New("Người dùng đã tồn tại")
	// ErrSignUpFailed err code
	ErrSignUpFailed = errors.New("Đăng kí thất bại")
)
