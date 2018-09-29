package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDb(dns string) error {
	var err error
	DB, err = sqlx.Open("mysql", dns)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(1000)
	DB.SetMaxIdleConns(16)
	return nil
}
