package req

// SignUpReq : Map of data request server for SignUp
type SignUpReq struct {
	FullName string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
