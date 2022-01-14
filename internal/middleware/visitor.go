package middleware

import (
	"fmt"
	"github.com/3115826227/go-web-live/internal/application"
	"github.com/3115826227/go-web-live/internal/handle"
	"github.com/3115826227/go-web-live/internal/log"
	"github.com/3115826227/go-web-live/internal/utils"
	"github.com/gin-gonic/gin"
)

func Visitor(c *gin.Context) {
	var ip = c.ClientIP()
	user, visitorExist, err := application.GetVisitorByIp(c, ip)
	if err != nil {
		log.Logger.Error(err.Error())
		handle.FailedResp(c, handle.CodeInternalError)
		return
	}
	var accountId string
	var username string
	if !visitorExist {
		// 新的访问用户
		accountId = utils.GenerateSerialNumber()
		username = fmt.Sprintf("访客%v", accountId[:5])
		if err = application.AddVisitor(c, accountId, username, ip); err != nil {
			log.Logger.Error(err.Error())
			handle.FailedResp(c, handle.CodeInternalError)
			return
		}
	} else {
		accountId = user.AccountId
		username = user.Username
	}
	var userMeta = handle.UserMeta{
		AccountId: accountId,
		Username:  username,
	}
	c.Set(handle.GinContextKeyUserMeta, &userMeta)
	c.Next()
}
