package rsp

type User struct {
	AccountID string `json:"account_id"`
	Username  string `json:"username"`
}

type UserDetail struct {
	User
	OpenLiveRoom bool `json:"open_live_room"`
}

type UserDataResp struct {
	AccountId string `json:"account_id"`
	Username  string `json:"username"`
}

type LoginResult struct {
	UserInfo UserDataResp `json:"user_info"`
	Token    string       `json:"token"`
}

type LiveRoomOriginUser struct {
	AccountID             string `json:"account_id"`
	Username              string `json:"username"`
	Visitor               bool   `json:"visitor"`
	PermissionSendMessage bool   `json:"permission_send_message"`
}

type LiveRoomUser struct {
	Username string `json:"username"`
	Visitor  bool   `json:"visitor"`
}
