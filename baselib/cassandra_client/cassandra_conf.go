/* !!
 * File: cassandra_conf.go
 * File Created: Thursday, 27th May 2021 10:19:46 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 10:26:45 am
 
 */

package cassandra_client

import "time"

// CassandraConfig type;
type CassandraConfig struct {
	Name           string // for trace
	Environment    string
	Hosts          []string // data source name
	Port           int
	Keyspace       string
	Timeout        int // in seconds; example 30 as 30s
	ConnectTimeout int // in seconds; example 30 as 30s
	Username       string
	Password       string
}

const (
	KDefaultTimeout = 30 * time.Second
)
