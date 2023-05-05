

package sql_client

import (
	"fmt"
	"sync"

	"github.com/TrHung-297/fountain/baselib/g_log"
)

// sqlClientManager type;
type sqlClientManager struct {
	sqlClients sync.Map
}

var sqlClientsManagerInstance = &sqlClientManager{}

// default value env key is "MySQL";
// if configKeys was set, key env will be first value (not empty) of this;
func InstallSQLClientManager(configKeys ...string) {
	getConfigFromEnv(configKeys...)

	for _, config := range configs {
		client := NewSqlxDB(config)
		if client == nil {
			err := fmt.Errorf("InstallSQLClientsManager - NewSqlxDB {%v} error", config)
			g_log.V(1).WithError(err).Errorf("InstallSQLClientsManager - Error: %v", err)

			panic(err)
		}

		if config.Name == "" {
			err := fmt.Errorf("InstallSQLClientsManager - config error: config.Name is empty")
			g_log.V(1).WithError(err).Errorf("InstallSQLClientsManager - Error: %v", err)

			panic(err)
		}
		if val, ok := sqlClientsManagerInstance.sqlClients.Load(config.Name); ok {
			err := fmt.Errorf("InstallSQLClientsManager - config error: duplicated config.Name {%v}", val)
			g_log.V(1).WithError(err).Errorf("InstallSQLClientsManager - Error: %v", err)

			panic(err)
		}

		sqlClientsManagerInstance.sqlClients.Store(config.Name, client)
	}
}

// GetSQLClient type;
func GetSQLClient(dbName string) (client *SQLClient) {
	if val, ok := sqlClientsManagerInstance.sqlClients.Load(dbName); ok {
		if client, ok = val.(*SQLClient); ok {
			return
		}
	}

	g_log.V(3).Infof("GetSQLClient - Not found client: %s", dbName)
	return
}

// GetSQLClientManager type;
func GetSQLClientManager() sync.Map {
	return sqlClientsManagerInstance.sqlClients
}
