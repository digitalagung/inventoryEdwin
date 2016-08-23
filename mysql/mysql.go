package mysql

import (
	"database/sql"
	"gopkg.in/gorp.v1"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


func InitDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:toor@/inventorydb")
	if err != nil {
		log.Fatalln("sql.Open failed", err)
	}

	return &gorp.DbMap{Db: db}
}