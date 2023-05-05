

package redis_dao

import (
	"context"

	"github.com/TrHung-297/fountain/baselib/g_log"
	"github.com/TrHung-297/fountain/baselib/redis_client"
	"github.com/TrHung-297/fountain/proto/g_proto"
)

type OpenIDCacheDAO struct {
	client *redis_client.RedisPool
}

func NewOpenIDCacheDAO(redis *redis_client.RedisPool) *OpenIDCacheDAO {
	return &OpenIDCacheDAO{
		client: redis,
	}
}

func (dao *OpenIDCacheDAO) GetAccessIDByDeviceKind(sessionKey string, deviceKind int) (accessID string) {
	conn := dao.client.Get()
	if conn == nil {
		g_log.V(1).Info("OpenIDCacheDAO::GetAccessIDByDeviceKind - Error: Can not get connection")

		return ""
	}

	field := ""

	if deviceKind == g_proto.KDeviceKindPC {
		field = g_proto.KCacheUserAuthorizationTokenField
	} else if deviceKind == g_proto.KDeviceKindMobile {
		field = g_proto.KCacheUserAuthorizationMobileTokenField
	} else if deviceKind == g_proto.KDeviceKindWeb {
		field = g_proto.KCacheUserAuthorizationWebTokenField
	}
	ctx, cancel := context.WithTimeout(context.Background(), redis_client.KDefaultTimeout)
	defer cancel()

	accessID, err := conn.HGet(ctx, sessionKey, field).Result()
	if err != nil {
		g_log.V(1).WithError(err).Errorf("OpenIDCacheDAO::GetAccessIDByDeviceKind - for sessionKey: %s with field: %s, Error: %+v", sessionKey, field, err)
	}

	return accessID
}
