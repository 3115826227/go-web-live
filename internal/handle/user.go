package handle

import (
	"github.com/3115826227/go-web-live/internal/application"
	"github.com/3115826227/go-web-live/internal/config"
	"github.com/3115826227/go-web-live/internal/db/infrastructure/dbclient"
	"github.com/3115826227/go-web-live/internal/db/tables"
	"github.com/3115826227/go-web-live/internal/dtos"
	"github.com/3115826227/go-web-live/internal/errors"
	"github.com/3115826227/go-web-live/internal/handle/requests"
	"github.com/3115826227/go-web-live/internal/handle/rsp"
	"github.com/3115826227/go-web-live/internal/log"
	"github.com/3115826227/go-web-live/internal/utils"
	"github.com/gin-gonic/gin"
	"time"
)

// 用户注册
func UserRegister(c *gin.Context) {
	var req requests.UserRegisterReq
	req.Username = c.PostForm("username")
	req.LoginName = c.PostForm("login_name")
	req.Password = c.PostForm("password")
	if err := c.Bind(&req); err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, errors.CodeInvalidParamError)
		return
	}
	var userRegister = dtos.UserRegister{
		Username:  req.Username,
		LoginName: req.LoginName,
		Password:  req.Password,
	}
	if err := application.UserRegister(c, userRegister); err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, err.Code())
		return
	}
	SuccessResp(c, "", nil)
}

// 用户登录
func UserLogin(c *gin.Context) {
	var req requests.UserLoginReq
	req.LoginName = c.PostForm("login_name")
	req.Password = c.PostForm("password")
	if err := c.Bind(&req); err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, errors.CodeInvalidParamError)
		return
	}
	user, err := application.UserLogin(c, req.LoginName, req.Password)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, err.Code())
		return
	}
	token, err1 := utils.GenerateToken(user.AccountId, time.Now(), config.GetConfig().TokenSecret)
	if err1 != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, errors.CodeInternalError)
		return
	}
	var loginResult = rsp.LoginResult{
		UserInfo: rsp.UserDataResp{
			AccountId: user.AccountId,
			Username:  user.Username,
		},
		Token: token,
	}
	go func() {
		var userMeta = &UserMeta{
			AccountId: user.AccountId,
			Username:  user.Username,
		}
		var userToken = tables.UserToken{
			Token:     token,
			AccountId: user.AccountId,
			Value:     userMeta.ToString(),
		}
		if err1 = dbclient.GetDBClient().AddUserToken(userToken); err != nil {
			log.Logger.Error(err.Error())
		}
	}()
	SuccessResp(c, "", loginResult)
}

// 用户获取信息
func UserDetail(c *gin.Context) {
	userMeta := GetUserMeta(c)
	user, err := application.GetUserById(c, userMeta.AccountId)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	var exist bool
	_, exist, err = application.GetLiveRoomByAccountId(c, userMeta.AccountId)
	if err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	var response = rsp.UserDetail{
		User: rsp.User{
			AccountID: user.AccountId,
			Username:  user.Username,
		},
		OpenLiveRoom: exist,
	}
	SuccessResp(c, "", response)
}

// 用户退出登录
func UserLogout(c *gin.Context) {
	userMeta := GetUserMeta(c)
	if err := dbclient.GetDBClient().DeleteUserTokenByAccountId(userMeta.AccountId); err != nil {
		log.Logger.Error(err.Error())
		FailedResp(c, CodeInternalError)
		return
	}
	SuccessResp(c, "", nil)
}
