

package dao

import (
	"sync"

	"github.com/TrHung-297/fountain/baselib/g_log"
	"github.com/TrHung-297/fountain/baselib/redis_client"
	"github.com/TrHung-297/fountain/biz/dal/dao/redis_dao"

	elastic "github.com/elastic/go-elasticsearch/v7"
	"github.com/jmoiron/sqlx"
)

// DB_MASTER const
const (
	DB_MASTER = "immaster"
	DB_SLAVE  = "imslave"

	GLOBAL_CACHE  = "gcache"
	CACHE         = "cache"
	SEQ           = "seq"
	OPEN_ID_CACHE = "open_id_cache"

	CASSYNC = "cassync"

	ELASTIC = "primary"
)

// ----------------------- MySQL ------------------------

type MysqlDAOList struct{}

// MysqlDAOManager type
type MysqlDAOManager struct {
	daoListMap map[string]*MysqlDAOList
}

var mysqlDAOManager = &MysqlDAOManager{make(map[string]*MysqlDAOList)}

// InstallMysqlDAOManager func
func InstallMysqlDAOManager(clients sync.Map) { /*map[string]*sqlx.DB*/
	clients.Range(func(key, value interface{}) bool {
		k, _ := key.(string)
		v, _ := value.(*sqlx.DB)

		daoList := &MysqlDAOList{}

		_ = v
		// daoList.CommonDAO = mysql_dao.NewCommonDAO(v)

		mysqlDAOManager.daoListMap[k] = daoList
		return true
	})
}

// GetMysqlDAOListMap func
func GetMysqlDAOListMap() map[string]*MysqlDAOList {
	return mysqlDAOManager.daoListMap
}

// GetMysqlDAOList func
func GetMysqlDAOList(dbName string) (daoList *MysqlDAOList) {
	daoList, ok := mysqlDAOManager.daoListMap[dbName]
	if !ok {
		g_log.V(1).Errorf("GetMysqlDAOList - Not found daoList: %s", dbName)
	}

	return
}

// GetCommonDAO func
// func GetCommonDAO(dbName string) (dao *mysql_dao.CommonDAO) {
// 	daoList := GetMysqlDAOList(dbName)
// 	if daoList != nil {
// 		dao = daoList.CommonDAO
// 	}
// 	return
// }

// ----------------------- Redis ------------------------
// RedisDAOList type
type RedisDAOList struct {
	OpenIDCacheDAO *redis_dao.OpenIDCacheDAO
}

// RedisDAOManager type
type RedisDAOManager struct {
	daoListMap map[string]*RedisDAOList
}

var redisDAOManager = &RedisDAOManager{make(map[string]*RedisDAOList)}

// InstallRedisDAOManager type
func InstallRedisDAOManager(clients map[string]*redis_client.RedisPool) {
	for k, v := range clients {
		daoList := &RedisDAOList{}

		daoList.OpenIDCacheDAO = redis_dao.NewOpenIDCacheDAO(v)

		redisDAOManager.daoListMap[k] = daoList
	}
}

// GetRedisDAOList type
func GetRedisDAOList(redisName string) (daoList *RedisDAOList) {
	daoList, ok := redisDAOManager.daoListMap[redisName]
	if !ok {
		g_log.V(1).Errorf("GetRedisDAOList - Not found daoList: %s", redisName)
	}
	return
}

// GetRedisDAOListMap type
func GetRedisDAOListMap() map[string]*RedisDAOList {
	return redisDAOManager.daoListMap
}

// GetRedisOpenIDCacheDAO func
func GetRedisOpenIDCacheDAO(redisName string) (dao *redis_dao.OpenIDCacheDAO) {
	daoList := GetRedisDAOList(redisName)
	if daoList != nil {
		dao = daoList.OpenIDCacheDAO
	}
	return
}

// ----------------------- Cassandra ------------------------

// ----------------------- Elastic ------------------------

// ElasticDAOList type
type ElasticDAOList struct {
}

// ElasticDAOManager type
type ElasticDAOManager struct {
	daoListMap map[string]*ElasticDAOList
}

var elasticDAOManager = &ElasticDAOManager{make(map[string]*ElasticDAOList)}

// InstallElasticDAOManager type
func InstallElasticDAOManager(clients sync.Map) {
	clients.Range(func(key, value interface{}) bool {
		k, _ := key.(string)
		v, _ := value.(*elastic.Client)

		_ = v
		_ = k
		daoList := &ElasticDAOList{}

		elasticDAOManager.daoListMap[k] = daoList

		return true
	})
}

// GetElasticDAOList type
func GetElasticDAOList(elasticName string) (daoList *ElasticDAOList) {
	daoList, ok := elasticDAOManager.daoListMap[elasticName]
	if !ok {
		g_log.V(1).Errorf("GetElasticDAOList - Not found daoList: %s", elasticName)
	}

	return
}

// GetElasticDAOListMap type
func GetElasticDAOListMap() map[string]*ElasticDAOList {
	return elasticDAOManager.daoListMap
}
