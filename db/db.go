package db

import (
	"database/sql"
	"todolist/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {

	conf := config.GetConfig()
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@/" + conf.DB_DATABASE

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}
}

func CreateConn() *sql.DB {
	return db
}
