package frame

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"wwbtg/com/db"

	_ "github.com/go-sql-driver/mysql"
)

// -- 为应用模块的model提供公共操作
type DefaultModel struct {
	db     *sql.DB
	TbName string
}

func (self *DefaultModel) InitDB(tb string) {
	self.db = db.GetMysqlDB()
	self.TbName = tb
}

func (self *DefaultModel) Select(fields, conds []string) *sql.Rows {
	oFields := strings.Join(fields, ",")

	oConds := strings.Join(conds, " and ")

	sql := fmt.Sprintf("select %s from %s where %s", oFields, self.TbName, oConds)
	log.Println("sql: " + sql)

	rows, err := self.db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}

	return rows
}
