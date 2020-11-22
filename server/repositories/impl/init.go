package impl

import (
	"gungnir/config"
	log "gungnir/logs"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	conn    *xorm.Engine
	connErr error
	once    sync.Once
)

func getDatabaseConn() (*xorm.Engine, error) {
	once.Do(func() {
		log.Debugln(config.GetConfig().GetString(config.DATABASE_CONNENTION_SOURCE))
		conn, connErr = xorm.NewEngine("mysql", config.GetConfig().GetString(config.DATABASE_CONNENTION_SOURCE))
		if connErr != nil {
			log.Errorln(connErr)
			return
		}
		conn.ShowSQL(true)
	})

	return conn, connErr
}
