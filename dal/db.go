package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB(dsn string) (err error) {
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(500)
	DB.SetMaxIdleConns(50)
	return
}
