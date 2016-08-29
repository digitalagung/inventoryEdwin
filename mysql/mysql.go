package mysql

import (
	"database/sql"
	"gopkg.in/gorp.v1"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/edwinlab/inventory/util"
)


func InitDb() *gorp.DbMap {
	db, err := sql.Open("mysql", util.GetEnv("MYSQL_URI"))
	if err != nil {
		log.Fatalln("sql.Open failed", err)
	}

	return &gorp.DbMap{Db: db}
}