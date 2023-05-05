/* !!
 * File: client_manager.go
 * File Created: Thursday, 20th May 2021 10:33:24 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 20th May 2021 10:34:53 am
 
 */

package elastic_client

import (
	"fmt"
	"sync"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"

	elastic "github.com/elastic/go-elasticsearch/v7"
)

// ElasticClientManager type;
type ElasticClientManager struct {
	elasticClients sync.Map
}

var elasticClients = &ElasticClientManager{}

// default value env key is "Elastic";
// if configKeys was set, key env will be first value (not empty) of this;
func InstallElasticClientManager(configKeys ...string) {
	getConfigFromEnv(configKeys...)

	if esConf == nil {
		err := fmt.Errorf("need config for elastic client first")
		panic(err)
	}

	client := NewElasticClient()
	if client == nil {
		err := fmt.Errorf("InstallElasticClientManager - NewElasticClient error")
		panic(err)
	}

	if val, ok := elasticClients.elasticClients.Load(esConf.Name); ok {
		g_log.V(1).Infof("InstallElasticClientManager - config error: duplicated config.Name %s with %v", esConf.Name, val)

		return
	}

	elasticClients.elasticClients.Store(esConf.Name, client)
}

// GetElasticClient type;
func GetElasticClient(dbName string) (client *elastic.Client) {
	if val, ok := elasticClients.elasticClients.Load(dbName); ok {
		if client, ok = val.(*elastic.Client); ok {
			return
		}
	}

	g_log.V(1).Infof("GetElasticClient - Not found client: %s", dbName)
	return
}

// GetElasticClientManager type;
func GetElasticClientManager() sync.Map {
	return elasticClients.elasticClients
}
