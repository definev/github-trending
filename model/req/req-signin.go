package req

// SignInReq : Map of data request server for SignUp
type SignInReq struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
