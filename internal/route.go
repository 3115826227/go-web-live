package internal

import (
	"github.com/3115826227/go-web-live/internal/handle"
	"github.com/3115826227/go-web-live/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(engine *gin.Engine) {
	engine.Static("static", "static")
	engine.LoadHTMLGlob("views/*")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	app := engine.Group("/api/v1")
	// 公共操作
	app.POST("/user/register", handle.UserRegister)
	app.POST("/user/login", handle.UserLogin)
	app.GET("/live/live_rooms", handle.LiveHandle)
	app.GET("/session", handle.SessionHandle)
	// 访客操作
	visitor := engine.Group("/api/v1", middleware.Visitor)
	visitor.GET("/live/live_room/detail/visitor", handle.LiveDetailVisitorQueryHandle)
	visitor.POST("/visitor/opt", handle.VisitorOperatorHandle)
	visitor.GET("/messages/visitor", handle.MessageVisitorQueryHandle)
	visitor.GET("/session/detail/visitor", handle.SessionDetailVisitorQueryHandle)

	// 用户操作
	app.Use(middleware.CheckToken)
	app.GET("/user/detail", handle.UserDetail)
	app.GET("/user/logout", handle.UserLogout)
	app.GET("/live/live_room", handle.LiveOriginHandle)
	app.GET("/live/live_room/detail/user", handle.LiveDetailUserQueryHandle)
	app.POST("/live/live_room/open", handle.OpenLiveHandle)
	app.GET("/live/live_room/users", handle.LiveUserHandle)
	app.POST("/live/live_room/origin/opt", handle.LiveOriginOperatorHandle)
	app.POST("/user/opt", handle.UserOperatorHandle)
	app.POST("/opt/user", handle.OperatorUserHandle)
	app.GET("/messages/user", handle.MessageUserQueryHandle)
	app.GET("/live/live_room/messages/origin", handle.LiveMessageOriginQueryHandle)
	app.POST("/live/live_room/message", handle.SendLiveMessageHandle)
}
