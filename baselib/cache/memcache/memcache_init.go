/* !!
 * File: memcache_init.go
 * File Created: Thursday, 27th May 2021 10:19:43 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 10:25:55 am
 
 */

package memcache

import (
	"encoding/json"
	"fmt"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/cache"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"

	"github.com/bradfitz/gomemcache/memcache"
)

// StartAndGC start memcache adapter.
// config string is like {"conn":"connection info"}.
// if connecting error, return.
func (mc *Cache) StartAndGC(configs ...string) error {
	var cf *Config

	for _, config := range configs {
		err := json.Unmarshal([]byte(config), &cf)
		if config != "" && err != nil {
			continue
		}
	}

	if cf == nil {
		return fmt.Errorf("not found any config")
	}

	mc.connInfo = cf
	if mc.conn == nil {
		if err := mc.connectInit(); err != nil {
			return err
		}
	}
	return nil
}

// connect to memcache and keep the connection.
func (mc *Cache) connectInit() error {
	g_log.V(3).Infof("Init connection to memcache server at: %+v", mc.connInfo.Conn)
	mc.conn = memcache.New(mc.connInfo.Conn...)
	return nil
}

func init() {
	cache.Register("memcache", NewMemCache)
}
