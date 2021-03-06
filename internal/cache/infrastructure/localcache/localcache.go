package localcache

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/3115826227/go-web-live/internal/cache/interfaces"
	"github.com/patrickmn/go-cache"
	"strings"
	"time"
)

var defaultTime = 24 * 60 * time.Minute

var (
	Nil                   = errors.New("cache:nil")
	KeyNull               = errors.New("key is null")
	ValueNull             = errors.New("value is null")
	FieldNull             = errors.New("field is null")
	HashSetFieldTypeError = errors.New("hash set field's type is not string")
	FieldValueNumberError = errors.New("hash set field and value number is fault")
)

type RamCacheClient struct {
	c *cache.Cache
}

func NewRamCacheClient() interfaces.Cache {
	return &RamCacheClient{
		c: cache.New(cache.DefaultExpiration, cache.DefaultExpiration),
	}
}

func (rcc *RamCacheClient) invalidKey(key string) error {
	if key == "" {
		return KeyNull
	}
	return nil
}

func (rcc *RamCacheClient) invalidField(field string) error {
	if field == "" {
		return FieldNull
	}
	return nil
}

func (rcc *RamCacheClient) invalidValue(value string) error {
	if value == "" {
		return ValueNull
	}
	return nil
}

func (rcc *RamCacheClient) invalidKeyValue(key, value string) error {
	if err := rcc.invalidKey(key); err != nil {
		return err
	}
	return rcc.invalidValue(value)
}

func (rcc *RamCacheClient) invalidFieldValue(field, value string) error {
	if err := rcc.invalidField(field); err != nil {
		return err
	}
	return rcc.invalidValue(value)
}

func (rcc *RamCacheClient) SetOneDay(key string, value interface{}) error {
	val, err := marshalValue(value)
	if err != nil {
		return err
	}
	if err = rcc.invalidKeyValue(key, val); err != nil {
		return err
	}
	rcc.c.Set(key, val, defaultTime)
	return nil
}

func (rcc *RamCacheClient) Set(key string, value interface{}, expiration time.Duration) error {
	val, err := marshalValue(value)
	if err != nil {
		return err
	}
	if err = rcc.invalidKeyValue(key, val); err != nil {
		return err
	}
	rcc.c.Set(key, val, expiration)
	return nil
}

func (rcc *RamCacheClient) Get(key string) (data string, err error) {
	value, exist := rcc.c.Get(key)
	if !exist {
		return "", Nil
	}
	data = value.(string)
	return
}

func (rcc *RamCacheClient) Del(keys ...string) error {
	for _, key := range keys {
		rcc.c.Delete(key)
	}
	return nil
}

func (rcc *RamCacheClient) getHashCache(key string) (*cache.Cache, bool, error) {
	if err := rcc.invalidKey(key); err != nil {
		return nil, false, err
	}
	data, exist := rcc.c.Get(key)
	if !exist {
		return nil, false, nil
	}
	value := data.(*cache.Cache)
	return value, true, nil
}

func (rcc *RamCacheClient) HSet(key string, values ...interface{}) error {
	hashCache, exist, err := rcc.getHashCache(key)
	if err != nil {
		return err
	}
	if !exist {
		hashCache = cache.New(cache.DefaultExpiration, cache.DefaultExpiration)
	}
	var field, val string
	for i := 0; i < len(values); i += 2 {
		// ??????panic????????????hash set field???????????????
		switch values[i].(type) {
		case string:
			field = values[i].(string)
		default:
			err = HashSetFieldTypeError
			return err
		}
		// hash set???field???value???????????????
		if i+1 >= len(values) {
			err = FieldValueNumberError
			return err
		}
		val, err = marshalValue(values[i+1])
		if err != nil {
			return err
		}
		field = fmt.Sprintf("%v:%v", key, field)
		if err = rcc.invalidFieldValue(field, val); err != nil {
			return err
		}
		hashCache.Set(field, val, cache.DefaultExpiration)
	}
	rcc.c.Set(key, hashCache, cache.DefaultExpiration)
	return nil
}

func (rcc *RamCacheClient) HGet(key, field string) (string, error) {
	hashCache, exist, err := rcc.getHashCache(key)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", Nil
	}
	var value interface{}
	value, exist = hashCache.Get(fmt.Sprintf("%v:%v", key, field))
	if !exist {
		return "", Nil
	}
	return value.(string), nil
}

func (rcc *RamCacheClient) HDel(key string, fields ...string) error {
	hashCache, exist, err := rcc.getHashCache(key)
	if err != nil {
		return err
	}
	if !exist {
		return nil
	}
	for _, field := range fields {
		hashCache.Delete(fmt.Sprintf("%v:%v", key, field))
	}
	return nil
}

func (rcc *RamCacheClient) HGetAll(key string) (map[string]string, error) {
	hashCache, exist, err := rcc.getHashCache(key)
	if err != nil {
		return nil, err
	}
	var mp = make(map[string]string)
	if !exist {
		return make(map[string]string), nil
	}
	for k, v := range hashCache.Items() {
		mp[strings.TrimPrefix(strings.TrimPrefix(k, key), ":")] = v.Object.(string)
	}
	return mp, nil
}

func (rcc *RamCacheClient) Close() {
	// nothing to do
}

func marshalValue(value interface{}) (string, error) {
	if value == nil {
		return "", ValueNull
	}
	switch value.(type) {
	case string:
		return value.(string), nil
	default:
		data, err := json.Marshal(value)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}
}
