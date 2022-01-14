package cache

import (
	"github.com/3115826227/go-web-live/internal/cache/infrastructure/localcache"
	"github.com/3115826227/go-web-live/internal/cache/interfaces"
)

var (
	cache interfaces.Cache
)

func GetCache() interfaces.Cache {
	return cache
}

func InitCache() {
	cache = localcache.NewRamCacheClient()
}
