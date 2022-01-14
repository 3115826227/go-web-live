package requests

import "github.com/3115826227/go-web-live/internal/constant"

type LiveOperatorReq struct {
	LiveOperator constant.LiveOperator `json:"live_operator" binding:"required"`
}

type LiveOriginOperatorReq struct {
	AccountId      string                  `json:"account_id"`
	BizType        constant.BizType        `json:"biz_type" binding:"required"`
	OriginOperator constant.OriginOperator `json:"origin_operator" binding:"required"`
}

type LiveUserOperatorReq struct {
	BizType      constant.BizType      `json:"biz_type" binding:"required"`
	BizId        string                `json:"biz_id" binding:"required"`
	UserOperator constant.UserOperator `json:"user_operator" binding:"required"`
}
