package db

import (
	"database/sql"
	"log"
	"sync"
	"time"
)

var instance *sql.DB
var once sync.Once

func GetMysqlDB() *sql.DB {
	once.Do(func() {
		dsn := "study:*@tcp(0.0.0.0:3389)/work?charset=utf8"

		var err error
		instance, err = sql.Open("mysql", dsn)

		if err != nil {
			log.Fatal(err)
		}

		instance.SetMaxOpenConns(20)
		instance.SetConnMaxLifetime(time.Minute * 60)

	})

	return instance
}
