package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var SQL *sqlx.DB

const dbName = "memo"

func DSN() string {
	return "root:123456@tcp(127.0.0.1:3306)/" + dbName + "?parseTime=true"
}

func Connect() {
	var err error
	if SQL, err = sqlx.Connect("mysql", DSN()); err != nil {
		log.Println("SQL Driver Error", err)
		return
	}
	if err = SQL.Ping(); err != nil {
		log.Println("Database Connect Fail", err)
	}
}
