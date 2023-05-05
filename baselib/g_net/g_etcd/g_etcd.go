package g_etcd

import (
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type GEtcd struct {
	*clientv3.Client
}

type EtcdLogLevel int

const (
	EtcdLogLevelDebug   = -2
	EtcdLogLevelInfo    = -1
	EtcdLogLevelWarning = 1
	EtcdLogLevelError   = 2
	EtcdLogLevelDPanic  = 3
	EtcdLogLevelPanic   = 4
	EtcdLogLevelFatal   = 5
)

type discoveryOption struct {
	EtcdAddrs      []string     `json:"etcd_addrs,omitempty"`
	EtcdUsername   string       `json:"etcd_username,omitempty"`
	EtcdPassword   string       `json:"etcd_password,omitempty"`
	GlobalInstance bool         `json:"global_instance,omitempty"`
	EtcdLogLevel   EtcdLogLevel `json:"log_level,omitempty"`
}

type DiscoveryOption func(*discoveryOption)

func WithEtcdAddrs(etcdAddr ...string) DiscoveryOption {
	return func(do *discoveryOption) {
		do.EtcdAddrs = etcdAddr
	}
}

func WithEtcdUsername(etcdUsername string) DiscoveryOption {
	return func(do *discoveryOption) {
		do.EtcdUsername = etcdUsername
	}
}

func WithEtcdPassword(etcdPwd string) DiscoveryOption {
	return func(do *discoveryOption) {
		do.EtcdPassword = etcdPwd
	}
}

// WithLogLevel func;
// -1 for debug, 0 for info, 1 for waning, 2 for error, 4 for d_panic, 5 for panic, 6 for fatal
func WithLogLevel(logLevel EtcdLogLevel) DiscoveryOption {
	return func(do *discoveryOption) {
		do.EtcdLogLevel = logLevel
	}
}

func WithIsGlobalInstance(isGlobalInstance bool) DiscoveryOption {
	return func(do *discoveryOption) {
		do.GlobalInstance = isGlobalInstance
	}
}

// InstanceEtcdManger func;
// If envKey is empty, set default for env key
func InstanceEtcdManger(envKey ...string) {
	if config == nil {
		createEtcdConfigFromEnv(envKey...)
	}

	if config == nil {
		err := fmt.Errorf("not found config for etcd")
		panic(err)
	}

	if config.LogLevel == 0 {
		config.LogLevel = EtcdLogLevelError
	}

	if config.LogLevel < 0 {
		config.LogLevel++
	}

	NewEtcdDiscovery(WithEtcdAddrs(config.Addrs...), WithEtcdUsername(config.Username), WithEtcdPassword(config.Password), WithIsGlobalInstance(true), WithLogLevel(config.LogLevel))
}

func NewEtcdDiscovery(opts ...DiscoveryOption) *GEtcd {
	options := &discoveryOption{
		EtcdLogLevel: 1,
	}
	for _, opt := range opts {
		opt(options)
	}

	client, err := clientv3.New(clientv3.Config{
		Endpoints: options.EtcdAddrs,
		Username:  options.EtcdUsername,
		Password:  options.EtcdPassword,
		LogConfig: &zap.Config{
			Level:    zap.NewAtomicLevelAt(zapcore.Level(options.EtcdLogLevel)),
			Encoding: "json",
		},
		DialTimeout: 30 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	i := &GEtcd{Client: client}
	if options.GlobalInstance {
		gEtcd = i
	}

	return i
}

var gEtcd *GEtcd

func GetEtcdDiscoveryInstance() *GEtcd {
	if gEtcd == nil {
		err := fmt.Errorf("need initialization etcd global instance connection first")
		panic(err)
	}

	return gEtcd
}
