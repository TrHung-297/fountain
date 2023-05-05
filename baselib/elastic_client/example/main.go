/* !!
 * File: main.go
 * File Created: Thursday, 20th May 2021 10:33:24 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 20th May 2021 10:34:44 am
 
 */

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"

	"github.com/spf13/viper"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/elastic_client"
)

func init() {
	viper.Set("Elastic.Name", "gtv_log")
	viper.Set("Elastic.Environment", "local")
	viper.Set("Elastic.Host", "localhost:9200")
	viper.Set("Elastic.Username", "elastic")
	viper.Set("Elastic.Password", "changeme")
}

func QueryData() {
	client := elastic_client.GetElasticClient("gtv_log")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.Count(
		client.Count.WithContext(ctx),
		client.Count.WithIndex("gtv_log_event*"),
	)

	if err != nil {
		err := fmt.Errorf("QueryData - Error is %q while expect nil", err.Error())
		panic(err)
	}

	if res.IsError() {
		err := fmt.Errorf("queryData - Response error is not empty")
		panic(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err := fmt.Errorf("QueryData - ReadAll response Error is %q while expect nil", err.Error())
		panic(err)
	}

	if len(data) == 0 {
		err := fmt.Errorf("queryData - ReadAll response is empty")
		panic(err)
	}

	g_log.V(5).Infof("Data: %s", data)
}

func main() {
	elastic_client.InstallElasticClientManager()
	QueryData()
}
