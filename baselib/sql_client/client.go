

package sql_client

import (
	"time"

	_ "github.com/lib/pq"

	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/jmoiron/sqlx"
	"github.com/TrHung-297/fountain/baselib/g_log"
)

type SQLClient struct {
	*sqlx.DB
}

func (client *SQLClient) Get() *sqlx.DB {
	return client.DB
}

// NewSqlxDB type;
// For MySQL, posgreSQL
func NewSqlxDB(c *SQLConfig) *SQLClient {
	db, err := sqlx.Connect(c.Driver, c.DSN)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("NewSqlxDB Connect db error: %+v", err)
	}

	db.SetMaxOpenConns(c.Active)
	db.SetMaxIdleConns(c.Idle)
	if c.Lifetime < 60 {
		c.Lifetime = 5 * 60
	}
	db.SetConnMaxLifetime(time.Duration(c.Lifetime) * time.Second)

	g_log.V(3).Infof("NewSqlxDB: %+v", db.Stats())

	return &SQLClient{db}
}
