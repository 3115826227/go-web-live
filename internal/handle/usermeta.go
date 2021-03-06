package handle

import (
	"encoding/json"
	"github.com/3115826227/go-web-live/internal/handle/rsp"
	"github.com/gin-gonic/gin"
)

const (
	HeaderUUID  = "requestID"
	HeaderToken = "token"
	TokenPrefix = "token"
	HeaderIP    = "IP"

	HeaderAccountId  = "accountId"
	HeaderUsername   = "username"
	HeaderSchoolId   = "schoolId"
	HeaderPlatform   = "platform"
	HeaderReqId      = "reqId"
	HeaderIsOfficial = "isOfficial"
	HeaderPhone      = "phone"

	GinContextKeyUserMeta = "userMeta"

	QueryId           = "id"
	QueryAccountId    = "account_id"
	QueryLikeUsername = "username"
	QueryLikeName     = "name"

	QueryPage     = "page"
	QueryPageSize = "page_size"
)

type UserMeta struct {
	//用户ID
	AccountId string `json:"accountId"`
	//用户名
	Username string `json:"username"`
	//请求ID
	ReqId string `json:"reqId"`
	//平台
	Platform string `json:"platform"`
	//是否为超级管理员
	IsOfficial bool `json:"isOfficial"`
}

func (meta *UserMeta) ToString() string {
	data, _ := json.Marshal(meta)
	return string(data)
}

func GetUserMeta(c *gin.Context) *UserMeta {
	return c.MustGet(GinContextKeyUserMeta).(*UserMeta)
}

func (meta *UserMeta) GetUser() rsp.User {
	return rsp.User{
		AccountID: meta.AccountId,
		Username:  meta.Username,
	}
}
