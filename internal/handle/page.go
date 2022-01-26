package handle

import (
	"fmt"
	"github.com/3115826227/go-web-live/internal/constant"
	"github.com/3115826227/go-web-live/internal/errors"
	"github.com/3115826227/go-web-live/internal/handle/requests"
	"github.com/gin-gonic/gin"
	"strconv"
)

func PageHandle(c *gin.Context) (requests.PageCommonReq, errors.Error) {
	var req requests.PageCommonReq
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = fmt.Sprintf("%v", constant.DefaultPage)
	}
	var page int
	var err error
	if page, err = strconv.Atoi(pageStr); err != nil {
		return requests.PageCommonReq{}, errors.NewCommonError(errors.CodeInvalidParamError)
	}
	if page <= 0 {
		page = constant.DefaultPage
	}
	pageSizeStr := c.Query("page_size")
	if pageSizeStr == "" {
		pageSizeStr = fmt.Sprintf("%v", constant.DefaultPageSize)
	}
	var pageSize int
	if pageSize, err = strconv.Atoi(pageSizeStr); err != nil {
		return requests.PageCommonReq{}, errors.NewCommonError(errors.CodeInvalidParamError)
	}
	if pageSize <= 0 {
		pageSize = constant.DefaultPageSize
	}
	req.Page = int64(page)
	req.PageSize = int64(pageSize)
	return req, nil
}
