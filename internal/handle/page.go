package handle

import (
	"fmt"
	"github.com/3115826227/go-web-live/internal/constant"
	"github.com/3115826227/go-web-live/internal/handle/requests"
	"github.com/gin-gonic/gin"
	"strconv"
)

func PageHandle(c *gin.Context) (req requests.PageCommonReq, err error) {
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = fmt.Sprintf("%v", constant.DefaultPage)
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return
	}
	if page <= 0 {
		page = constant.DefaultPage
	}
	pageSizeStr := c.Query("page_size")
	if pageSizeStr == "" {
		pageSizeStr = fmt.Sprintf("%v", constant.DefaultPageSize)
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		return
	}
	if pageSize <= 0 {
		pageSize = constant.DefaultPageSize
	}
	req.Page = int64(page)
	req.PageSize = int64(pageSize)
	return
}
