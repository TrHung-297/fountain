
package g_proto

import "time"

const (
	// AWS
	KDefaultAwsRegion     string = "ap-southeast-1"
	KDefaultAwsPermission string = "public-read"
)

// constant for open id
const (
	KCacheUserAuthorizationFormat           string = "backend-user-session:%s" // userUUID
	KCacheUserAuthorizationTokenField       string = "token"
	KCacheUserAuthorizationMobileTokenField string = "mobile-token"
	KCacheUserAuthorizationWebTokenField    string = "web-token"
)

// constant for expire time
const (
	KCacheExpiresInForever        = 0
	KCacheExpiresInOneMinute      = time.Minute
	KCacheExpiresInTenMinutes     = time.Minute * 10
	KCacheExpiresInFifteenMinutes = time.Minute * 15
	KCacheExpiresInOneHour        = time.Hour
	KCacheExpiresInTwoHour        = KCacheExpiresInOneHour * 2
	KCacheExpiresInOneDay         = KCacheExpiresInOneHour * 24
	KCacheExpiresInThreeDay       = KCacheExpiresInOneDay * 3
	KCacheExpiresInOneMonth       = KCacheExpiresInOneDay * 30
	KCacheExpiresInThreeMonths    = KCacheExpiresInOneMonth * 3
	KCacheExpiresInSixMonths      = KCacheExpiresInOneMonth * 6
	KCacheExpiresInOneYear        = KCacheExpiresInOneMonth * 12
)

const (
	KDeviceKindPC     int = 1
	KDeviceKindMobile int = 3
	KDeviceKindWeb    int = 5
)
