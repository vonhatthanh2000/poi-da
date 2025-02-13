package global

import (
	"database/sql"
	"layerg-poi-da/pkg/logger"
	"layerg-poi-da/pkg/setting"

	dbConn "layerg-poi-da/db/sqlc"

	"github.com/redis/go-redis/v9"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Cdb    *sql.DB
	Rdb    *redis.Client
	Query  *dbConn.Queries
)
