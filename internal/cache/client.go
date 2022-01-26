package cache

import (
	"github.com/3115826227/go-web-live/internal/cache/infrastructure/localcache"
	"github.com/3115826227/go-web-live/internal/cache/interfaces"
	"github.com/3115826227/go-web-live/internal/db/tables"
	"github.com/3115826227/go-web-live/internal/errors"
)

type Client struct {
	cache interfaces.Cache
}

var (
	client interfaces.CacheClient
)

func GetCache() interfaces.CacheClient {
	return client
}

func InitCache() {
	client = &Client{cache: localcache.NewRamCacheClient()}
}

func (c *Client) SetLiveRoom(liveRoom tables.LiveRoom) errors.Error {
	return nil
}

func (c *Client) GetLiveRoomById(liveRoomId string) (tables.LiveRoom, errors.Error) {
	return tables.LiveRoom{}, nil
}

func (c *Client) DeleteLiveRoomById(liveRoomId string) errors.Error {
	return nil
}
