/* !!
 * File: cassandra_client.go
 * File Created: Thursday, 27th May 2021 10:19:46 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 10:26:41 am
 
 */

package cassandra_client

import (
	"time"

	"github.com/gocql/gocql"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"
)

// NewCassandraSession type;
func NewCassandraSession(c *CassandraConfig) *gocql.Session {
	cluster := gocql.NewCluster(c.Hosts...)

	cluster.Port = c.Port
	cluster.Consistency = gocql.Quorum
	cluster.Keyspace = c.Keyspace
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Duration(c.ConnectTimeout)
	cluster.Timeout = time.Duration(c.Timeout)
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: c.Username, Password: c.Password}
	cluster.Events.DisableNodeStatusEvents = false
	cluster.Events.DisableSchemaEvents = false
	cluster.Events.DisableTopologyEvents = false

	session, err := cluster.CreateSession()
	if err != nil {
		g_log.V(1).WithError(err).Errorf("NewCassandraSession - Error: %+v", err)
		return nil
	}

	return session
}
