/* !!
 * File: send_log_test.go
 * File Created: Wednesday, 5th May 2021 3:20:12 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Wednesday, 5th May 2021 3:20:12 pm
 
 */

package gostash

import (
	"testing"

	"github.com/spf13/viper"
)

func init() {
	viper.Set("logstash.Hosts", "ec2-54-169-231-186.ap-southeast-1.compute.amazonaws.com:9600")
	viper.Set("logstash.AppKey", "gtv_test")
	viper.Set("logstash.SecretKey", "log")
}

// Test that TestLogStashInsert works
func TestLogStashInsert(t *testing.T) {
	instance := InstallLogStashClient()
	if instance == nil {
		t.Errorf("TestLogStashInsert - Error can not create logstash instance")
	}

	dataLog := MakeElasticMessage()
	dataLog.DataLog.SetEventName("TestEvent")
	dataLog.DataLog.SetLogDataJSON(`{"Data":"This is data log"`)

	instance.InsertLog(dataLog)
}

func BenchmarkHeader(b *testing.B) {
	instance := InstallLogStashClient()
	if instance == nil {
		b.Errorf("TestLogStashConsumer - Error can not create logstash instance")
	}

	dataLog := MakeElasticMessage()
	dataLog.DataLog.SetEventName("TestEvent")
	dataLog.DataLog.SetLogDataJSON(`{"Data":"This is data log"`)

	for i := 0; i < b.N; i++ {
		instance.InsertLogSync(dataLog)
	}
}
