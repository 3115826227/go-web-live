package requests

type UserRegisterReq struct {
	Username  string `json:"username" binding:"required"`
	LoginName string `json:"login_name"  binding:"required"`
	Password  string `json:"password"  binding:"required"`
}

type UserLoginReq struct {
	LoginName string `json:"login_name"  binding:"required"`
	Password  string `json:"password"  binding:"required"`
}
