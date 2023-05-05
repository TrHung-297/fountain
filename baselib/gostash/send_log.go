/* !!
 * File: send_log.go
 * File Created: Wednesday, 5th May 2021 10:37:21 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Wednesday, 5th May 2021 3:20:01 pm
 
 */

package gostash

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/TrHung-297/fountain/baselib/base"
	"github.com/TrHung-297/fountain/baselib/g_log"
)

// ElasticConfig type
type ElasticConfig struct {
	LogAddr         []string
	AppKey          string
	SecretKey       string
	PlatformDefault string
	RouterDefault   string
}

//SetAppKey func
func SetAppKey(newAppKey string) string {
	if newAppKey != "" {
		logstashConf.AppKey = newAppKey
	}
	return logstashConf.AppKey
}

// SetSecretKey func
func SetSecretKey(newSecretKey string) string {
	if newSecretKey != "" {
		logstashConf.SecretKey = newSecretKey
	}
	return logstashConf.SecretKey
}

// ElasticMessage type
type ElasticMessage struct {
	TimeStamp string      `json:"@timestamp"`
	Index     string      `json:"__index"`
	LogType   string      `json:"__logtype"`
	AppKey    string      `json:"appkey"`
	SecretKey string      `json:"secret_key"`
	Platform  string      `json:"platform"`
	DataLog   DataLogBase `json:"data_log"`
}

// MakeElasticMessage func
func MakeElasticMessage() ElasticMessage {
	return ElasticMessage{
		AppKey:    logstashConf.AppKey,
		SecretKey: logstashConf.SecretKey,
		Platform:  logstashConf.PlatformDefault,
		LogType:   logstashConf.LogTypeDefault,
		DataLog:   NewDefaultDataLog(),
	}
}

// ToString func
func (e ElasticMessage) ToString() string {
	return base.JSONDebugDataString(e)
}

// LogStashClient type
type LogStashClient struct {
	clients  []*Gostash
	curIndex int
}

var (
	logStashClient *LogStashClient
)

// InstallLogStashClient func
func InstallLogStashClient() *LogStashClient {
	createConfigFromEnv()

	count := 0
	for logStashClient == nil && count < 5 {
		count++
		myG := NewList(logstashConf.Hosts, int(math.MaxInt32))
		for _, g := range myG {
			conn, err := g.Connect()
			if err != nil {
				err := fmt.Errorf("InstallLogStashClient - Can not create connection to %s:%d - Error: %+v", g.Hostname, g.Port, err)
				g_log.V(1).WithError(err).Errorf("Gostash::InstallLogStashClient - Error: %+v", err)

				panic(err)
			}

			g.Connection = conn
			g_log.V(3).Infof("Connect(): conn: %+v ", conn)
		}

		logStashClient = &LogStashClient{
			clients:  myG,
			curIndex: 0,
		}

		time.Sleep(500 * time.Millisecond)
	}

	if count >= 5 {
		err := fmt.Errorf("installLogStashClient - Can not create logstash's instace")
		panic(err)
	}

	return logStashClient
}

// GetLogStashClient func
func GetLogStashClient() *LogStashClient {
	if logStashClient == nil {
		return InstallLogStashClient()
	}

	return logStashClient
}

// GetNextClient func;
func (e *LogStashClient) GetNextClient() int {
	e.curIndex++
	if e.curIndex >= len(e.clients) {
		e.curIndex = 0
	}

	return e.curIndex
}

///////////////////////////////////////////
//tam thoi
func random() int {
	return rand.Intn(1)
}

// InsertLog func
func (e *LogStashClient) InsertLog(dataLogs ElasticMessage) {
	g_log.V(5).Infof("LogStashClient::InsertLog - logStashClient Addr: %p - With list conn: %+v", e, e.clients)

	go e.InsertLogSync(dataLogs)
}

// InsertLog func
func (e *LogStashClient) InsertLogSync(dataLogs ElasticMessage) {
	g_log.V(5).Infof("LogStashClient::insertLogSync - logStashClient Addr: %p - With list conn: %+v", e, e.clients)

	dataLogs.Index = strings.ToLower(fmt.Sprintf("%s.%s-%s", dataLogs.AppKey, dataLogs.Platform, dataLogs.LogType))
	dataLogs.TimeStamp = base.GetTimestampData()
	dataLogs.LogType = "log"

	retries := 0
	for retries < 3 {
		if len(e.clients) == 0 {
			g_log.V(1).Infof("LogStashClient::insertLogSync - No connection!")
			time.Sleep(1 * time.Second)
			retries++

			continue
		}

		rand.Seed(time.Now().UnixNano())
		myIdx := random()
		// myIdx := e.GetNextClient()
		myConn := e.clients[myIdx]
		err := myConn.Writeln(dataLogs.ToString())

		if err != nil {
			g_log.V(1).WithError(err).Errorf("LogStashClient::insertLogSync - Write to %v - err: %+v", myConn.Connection, err)
			g_log.V(1).WithError(err).Errorf("LogStashClient::insertLogSync - InsertLog into [%s] was error: %+v", myConn.String(), err)
			retries++
			if retries > 3 {
				return
			}

			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}

	if retries >= 3 {
		g_log.V(1).Infof("LogStashClient::insertLogSync - Retries too much")
	}
}
