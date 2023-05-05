/* !!
 * File: main.go
 * File Created: Thursday, 20th May 2021 10:33:24 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 20th May 2021 10:36:47 am
 
 */

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/spf13/viper"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/base"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/gostash"
)

func init() {
	// viper.Set("logstash.Hosts", "ec2-54-169-231-186.ap-southeast-1.compute.amazonaws.com:5044")
	viper.Set("LogStash.Hosts", "localhost:5044")
	viper.Set("LogStash.AppKey", "gtv_log_event")
	viper.Set("LogStash.SecretKey", "log")
}

func main() {
	instance := gostash.InstallLogStashClient()
	if instance == nil {
		err := fmt.Errorf("testLogStashInsert - Error can not create logstash instance")
		panic(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		for {
			dataLog := gostash.MakeElasticMessage()

			logMsg := gostash.NewDefaultDataLog()
			logMsg.EventName = "Install"

			uid, err := uuid.NewV4()
			if err != nil {
				logMsg.LogDataJSON = fmt.Sprintf(`{"Data":"This is data log: %s"`, err.Error())
			} else {
				logMsg.LogDataJSON = fmt.Sprintf(`{"Data":"This is data log: %s"`, uid.String())
			}

			dataLog.DataLog = logMsg

			log.Printf("dataLog: %s", base.JSONDebugDataString(dataLog))
			instance.InsertLogSync(dataLog)

			time.Sleep(500 * time.Millisecond)
		}
	}()

	<-ch
}
