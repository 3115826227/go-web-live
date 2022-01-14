package requests

import "github.com/3115826227/go-web-live/internal/constant"

type PageCommonReq struct {
	PageSize int64 `json:"page_size"`
	Page     int64 `json:"page"`
}

func (req *PageCommonReq) Validate() {
	if req.PageSize <= 0 {
		req.PageSize = constant.DefaultPageSize
	}
	if req.Page <= 0 {
		req.Page = constant.DefaultPage
	}
}
