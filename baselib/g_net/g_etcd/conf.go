package g_etcd

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"github.com/TrHung-297/fountain/baselib/g_log"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// NameProjectDir = "/gtv_plus"
const (
	KNameProjectDir = "/gtv_plus"
)

// ErrorWatcherClosed var
var ErrorWatcherClosed = fmt.Errorf("naming: watch closed")

type Config struct {
	Addrs    []string     `json:"addrs"`
	Username string       `json:"username"`
	Password string       `json:"password"`
	LogLevel EtcdLogLevel `json:"log_level"`
}

// Option type
type Option struct {
	Config      clientv3.Config
	RegistryDir string
	ServiceName string
	ServerID    string
	NodeData    NodeData
	TTL         time.Duration
}

// NodeData type
type NodeData struct {
	Addr     string
	Metadata interface{}
}

var config *Config

func createEtcdConfigFromEnv(envKey ...string) {
	config = &Config{}

	key := "Etcd"
	if len(envKey) > 0 {
		key = envKey[0]
	}

	if err := viper.UnmarshalKey(key, config); err != nil {
		g_log.V(1).Errorf("createEtcdConfigFromEnv - Error: %v from key %q", err, key)
		panic(err)
	}

	if len(config.Addrs) == 0 {
		err := fmt.Errorf("not found any addr as host for etcd at %q", fmt.Sprintf("%s.addrs", key))
		panic(err)
	}

}
