/* !!
 * File: client_manager_test.go
 * File Created: Thursday, 20th May 2021 10:33:24 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 20th May 2021 10:34:57 am
 
 */

package elastic_client

import (
	"context"
	"io/ioutil"
	"log"
	"testing"
	"time"

	elastic "github.com/elastic/go-elasticsearch/v7"
	"github.com/spf13/viper"
)

func init() {
	viper.Set("Elastic.Name", "gtv_log")
	viper.Set("Elastic.Environment", "local")
	viper.Set("Elastic.Host", "localhost:9200")
	viper.Set("Elastic.Username", "elastic")
	viper.Set("Elastic.Password", "changeme")
}

func TestClientManager(t *testing.T) {
	InstallElasticClientManager()

	mapClient := GetElasticClientManager()
	mapClient.Range(func(k, v interface{}) bool {
		if key, ok := k.(string); !ok {
			t.Errorf("TestClientManager - Can not parse name of instance")
		} else {
			if key != "gtv_log" {
				t.Errorf("TestClientManager - Name of instance is %q while expect %q", key, "gtv_log")
			}
		}

		if instance, ok := v.(*elastic.Client); !ok {
			t.Errorf("TestClientManager - Can not parse instance of Elastic")
		} else if instance == nil {
			t.Errorf("TestClientManager - Instance of Elastic is Nil")
		}

		return true
	})
}

func TestClientQueryManager(t *testing.T) {
	InstallElasticClientManager()

	mapClient := GetElasticClientManager()
	mapClient.Range(func(k, v interface{}) bool {
		if key, ok := k.(string); !ok {
			t.Errorf("TestClientManager - Can not parse name of instance")
		} else {
			if key != "gtv_log" {
				t.Errorf("TestClientManager - Name of instance is %q while expect %q", key, "gtv_log")
			}
		}

		if instance, ok := v.(*elastic.Client); !ok {
			t.Errorf("TestClientManager - Can not parse instance of Elastic")
		} else if instance == nil {
			t.Errorf("TestClientManager - Instance of Elastic is Nil")
		}

		return true
	})
}

func QueryData(t testing.TB) {
	client := GetElasticClient("gtv_log")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.Count(
		client.Count.WithContext(ctx),
		client.Count.WithIndex("gtv_log_event*"),
	)

	if err != nil {
		t.Errorf("QueryData - Error is %q while expect nil", err.Error())
	}

	if res.IsError() {
		t.Errorf("QueryData - Response error is not empty")
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("QueryData - ReadAll response Error is %q while expect nil", err.Error())
	}

	if len(data) == 0 {
		t.Errorf("QueryData - ReadAll response is empty")
	}

	log.Printf("Data: %s", data)
}

func TestClientQuery(t *testing.T) {
	InstallElasticClientManager()
	QueryData(t)
}

func BenchmarkClientQuery(b *testing.B) {
	InstallElasticClientManager()
	for i := 0; i < b.N; i++ {
		QueryData(b)
	}
}
