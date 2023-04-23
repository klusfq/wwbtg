package frame

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func getDB() *sql.DB {
	dsn := "study:future@tcp(121.37.71.178:3389)/work/charset=utf8"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(time.Minute * 60)

	return db
}

// -- 为应用模块的model提供公共操作
type DefaultModel struct {
	db     *sql.DB
	TbName string
}

func (self *DefaultModel) InitDB(tb string) {
	self.db = getDB()
	self.TbName = tb
}

func (self *DefaultModel) Select(fields, conds []string) *sql.Rows {
	oFields := strings.Join(fields, ",")

	oConds := strings.Join(conds, ",")

	sql := fmt.Sprintf("select %s from %s where %s", oFields, self.TbName, oConds)

	rows, err := self.db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}

	return rows
}
