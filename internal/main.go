package internal

import (
	"fmt"
	"github.com/3115826227/go-web-live/internal/cache"
	"github.com/3115826227/go-web-live/internal/config"
	"github.com/3115826227/go-web-live/internal/db/infrastructure/dbclient"
	"github.com/3115826227/go-web-live/internal/log"
	"github.com/gin-gonic/gin"
)

var (
	conf       config.Config
	serverName = "web-live"
)

func init() {
	// 初始化配置文件并获取
	conf = config.GetConfig()
	if err := log.InitLog(serverName, conf.Log.LogLevel, conf.Log.LogPath); err != nil {
		panic(err)
	}
	log.Logger.Info("log init successful")
	if err := dbclient.InitDB(conf.Database.DataSource, log.Logger); err != nil {
		panic(err)
	}
	log.Logger.Info("database init successful")
	cache.InitCache()
	log.Logger.Info("cache init successful")
}

func Main() {
	engine := gin.Default()
	Register(engine)
	if err := engine.Run(fmt.Sprintf("%v:%v", conf.Server.Host, conf.Server.Port)); err != nil {
		panic(err)
	}
}
