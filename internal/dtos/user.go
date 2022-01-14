package dtos

type UserRegister struct {
	Username  string
	LoginName string
	Password  string
}

type User struct {
	AccountId             string
	Username              string
	Visitor               bool
	PermissionSendMessage bool
}
