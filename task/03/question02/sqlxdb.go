package question02

import (
	"go-demo/task/03/db"
	"time"

	"github.com/jmoiron/sqlx"
)

func ConnectDB() *sqlx.DB {
	dsn := db.GetDsn()
	connect := sqlx.MustConnect("mysql", dsn)

	connect.SetMaxOpenConns(20)
	connect.SetMaxIdleConns(10)
	connect.SetConnMaxLifetime(30 * time.Minute)
	return connect
}
