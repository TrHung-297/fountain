/* !!
 * File: cassandra_client_manager.go
 * File Created: Thursday, 27th May 2021 10:19:46 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 10:26:37 am
 
 */

package cassandra_client

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"

	"github.com/gocql/gocql"
)

// CassandraClientManager type;
type CassandraClientManager struct {
	cassandraClients sync.Map
}

var cassandraManager = &CassandraClientManager{}

// InstallCassandraClientManager type;
func InstallCassandraClientManager(configs []CassandraConfig) {
	for _, config := range configs {
		session := NewCassandraSession(&config)
		if session == nil {
			err := fmt.Errorf("InstallCassandraClientManager - NewSqlxDB {%v} error", config)
			panic(err)
		}

		if config.Name == "" {
			err := fmt.Errorf("InstallCassandraClientManager - config error: config.Name is empty")
			panic(err)
		}
		if val, ok := cassandraManager.cassandraClients.Load(config.Name); ok {
			err := fmt.Errorf("InstallCassandraClientManager - config error: duplicated config.Name {%v}", val)
			panic(err)
		}
		cassandraManager.cassandraClients.Store(config.Name, session)

		go func() {
			var ch = make(chan os.Signal, 1)
			signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

			<-ch
			session.Close()
		}()
	}
}

// GetCassandraClient type;
func GetCassandraClient(dbName string) (session *gocql.Session) {
	if val, ok := cassandraManager.cassandraClients.Load(dbName); ok {
		if session, ok = val.(*gocql.Session); ok {
			return
		}
	}

	g_log.V(1).Infof("GetCassandraClient - Not found client: %s", dbName)
	return
}

// GetCassandraClientManager type;
func GetCassandraClientManager() sync.Map {
	return cassandraManager.cassandraClients
}
