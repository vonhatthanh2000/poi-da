package initialize

import (
	"database/sql"
	"layerg-poi-da/global"
	"time"

	dbConn "layerg-poi-da/db/sqlc"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func setPool(cdb *sql.DB) {
	m := global.Config.Database
	cdb.SetMaxOpenConns(m.MaxOpenConns)
	cdb.SetMaxIdleConns(m.MaxIdleConns)
	cdb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime) * time.Second)
}

func InitDB() {
	conn, err := sql.Open(
		global.Config.Database.Driver,
		global.Config.Database.Url,
	)

	if err != nil {
		checkErrorPanic(err, "Failed to connect to database")
	}

	global.Cdb = conn
	global.Query = dbConn.New(conn)
	global.Logger.Info("Connect database succesfully")

	// set pool
	setPool(global.Cdb)
}
