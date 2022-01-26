package middleware

import (
	"encoding/json"
	"github.com/3115826227/go-web-live/internal/db/infrastructure/dbclient"
	"github.com/3115826227/go-web-live/internal/errors"
	"github.com/3115826227/go-web-live/internal/handle"
	"github.com/3115826227/go-web-live/internal/handle/rsp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckToken(c *gin.Context) {
	var token = c.GetHeader(handle.HeaderToken)

	if token == "" {
		token = c.Query(handle.HeaderToken)
	}
	if token == "" {
		c.JSON(http.StatusOK, rsp.CommonResp{
			Code:    handle.CodeRequiredLogin,
			Message: handle.CodeRequiredLoginMsg,
			Data:    make([]interface{}, 0),
		})
		c.Abort()
		return
	}

	var userToken, err = dbclient.GetDBClient().GetUserTokenByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, rsp.CommonResp{
			Code:    err.Code(),
			Message: err.Error(),
			Data:    make([]interface{}, 0),
		})
		c.Abort()
		return
	}
	var str = userToken.Value
	var userMeta handle.UserMeta
	var err1 = json.Unmarshal([]byte(str), &userMeta)
	if err1 != nil {
		err = errors.NewCommonError(errors.CodeUnLoginError)
		c.JSON(http.StatusOK, rsp.CommonResp{
			Code:    err.Code(),
			Message: err.Error(),
			Data:    make([]interface{}, 0),
		})
		c.Abort()
		return
	}
	c.Set(handle.GinContextKeyUserMeta, &userMeta)
	c.Next()
}
