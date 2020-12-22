package errlog

import "errors"

var (
	// ErrUserExist err code
	ErrUserExist = errors.New("Người dùng đã tồn tại")
	// ErrUserNotFound err code
	ErrUserNotFound = errors.New("Người dùng không tồn tại")
	// ErrLogInFailed err code
	ErrLogInFailed = errors.New("Đăng nhập thất bại")
	// ErrSignUpFailed err code
	ErrSignUpFailed = errors.New("Đăng kí thất bại")
)
