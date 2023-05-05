/* !!
 * File: memcache.go
 * File Created: Thursday, 27th May 2021 10:19:43 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 10:26:22 am
 
 */

package memcache

import (
	"errors"
	"time"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/cache"

	"github.com/bradfitz/gomemcache/memcache"
)

// Config type
type Config struct {
	Conn []string `json:"conn"`
}

// Cache type
// Cache Memcache adapter.
type Cache struct {
	conn     *memcache.Client
	connInfo *Config
}

// NewMemCache create new memcache adapter.
func NewMemCache() cache.Cache {
	return &Cache{}
}

// Get get value from memcache.
func (mc *Cache) Get(key string) interface{} {
	if mc.conn == nil {
		if err := mc.connectInit(); err != nil {
			return nil
		}
	}
	if item, err := mc.conn.Get(key); err == nil {
		return item.Value
	}
	return nil
}

// GetMulti get value from memcache.
func (mc *Cache) GetMulti(keys []string) []interface{} {
	size := len(keys)
	var rv []interface{}
	if mc.conn == nil {
		if err := mc.connectInit(); err != nil {
			for i := 0; i < size; i++ {
				rv = append(rv, err)
			}
			return rv
		}
	}
	mv, err := mc.conn.GetMulti(keys)
	if err == nil {
		for _, v := range mv {
			rv = append(rv, v.Value)
		}
		return rv
	}
	for i := 0; i < size; i++ {
		rv = append(rv, err)
	}
	return rv
}

// Put put value to memcache.
func (mc *Cache) Put(key string, val interface{}, timeout time.Duration) error {
	if mc.conn == nil {
		if err := mc.connectInit(); err != nil {
			return err
		}
	}
	item := memcache.Item{Key: key, Expiration: int32(timeout / time.Second)}
	if v, ok := val.([]byte); ok {
		item.Value = v
	} else if str, ok := val.(string); ok {
		item.Value = []byte(str)
	} else {
		return errors.New("val only support string and []byte")
	}
	return mc.conn.Set(&item)
}

// PutWithoutExprise func;
func (mc *Cache) PutWithoutExprise(key string, val interface{}, timeout time.Duration) error {
	return mc.Put(key, val, 0)
}

// Delete delete value in memcache.
func (mc *Cache) Delete(key string) error {
	if mc.conn == nil {
		if err := mc.connectInit(); err != nil {
			return err
		}
	}
	return mc.conn.Delete(key)
}

// Incr increase counter.
func (mc *Cache) Incr(key string) error {
	if mc.conn == nil {
		if err := mc.connectInit(); err != nil {
			return err
		}
	}
	_, err := mc.conn.Increment(key, 1)
	return err
}

// Decr decrease counter.
func (mc *Cache) Decr(key string) error {
	if mc.conn == nil {
		if err := mc.connectInit(); err != nil {
			return err
		}
	}
	_, err := mc.conn.Decrement(key, 1)
	return err
}

// IsExist check value exists in memcache.
func (mc *Cache) IsExist(key string) bool {
	if mc.conn == nil {
		if err := mc.connectInit(); err != nil {
			return false
		}
	}
	_, err := mc.conn.Get(key)
	return err == nil
}

// ClearAll clear all cached in memcache.
func (mc *Cache) ClearAll() error {
	if mc.conn == nil {
		if err := mc.connectInit(); err != nil {
			return err
		}
	}
	return mc.conn.FlushAll()
}
