package req

// SignUpReq : Map of data request server for SignUp
type SignUpReq struct {
	FullName string `validate:"required" json:"fullName"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
