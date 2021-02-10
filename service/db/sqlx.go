package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var SqlxDB *sqlx.DB

func SetupSqlxDB(dbConfig DatabaseConfig) *sqlx.DB {
	SqlxDB = sqlx.MustConnect("postgres", ConnectionString(dbConfig))
	return SqlxDB
}
