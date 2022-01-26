package rsp

import "github.com/3115826227/go-web-live/internal/errors"

type CommonResp struct {
	Code    errors.ErrorCode `json:"code"`
	Message string           `json:"message"`
	Data    interface{}      `json:"data"`
}

type CommonListResp struct {
	List     []interface{} `json:"list"`
	Page     int64         `json:"page"`
	PageSize int64         `json:"page_size"`
	Total    int64         `json:"total"`
}
