package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

// Init init
func Init() (err error) {
	addr := "root:admin123@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	DB, err = sqlx.Connect("mysql", addr)
	if err != nil {
		return err
	}
	// max connection
	DB.SetMaxOpenConns(100)
	// max idle
	DB.SetMaxIdleConns(16)
	fmt.Println("Connect database succeed!")
	return
}
