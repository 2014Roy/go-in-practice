package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:123@/TESTDB?charset=utf8")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	SqlDB.SetMaxOpenConns(200)
	SqlDB.SetMaxIdleConns(100)
	fmt.Print("init mysql")
}
