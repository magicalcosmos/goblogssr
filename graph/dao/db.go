package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// Init init
func Init() (err error) {
	addr := "root:admin123@tcp(127.0.0.1:3306)/blogger"
	db, err = sqlx.Connect("mysql", addr)
	if err != nil {
		return err
	}
	// max connection
	db.SetMaxOpenConns(100)
	// max idle
	db.SetMaxIdleConns(16)
	return
}
