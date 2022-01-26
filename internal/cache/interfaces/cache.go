package interfaces

import "time"

type Cache interface {
	Close()
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (value string, err error)
	Del(key ...string) error
	HSet(key string, values ...interface{}) error
	HGet(key, field string) (string, error)
	HDel(key string, fields ...string) error
	HGetAll(key string) (map[string]string, error)
}

type CacheClient interface {

}