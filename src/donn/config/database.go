package config

import (
	"database/sql"
	"donn/entity"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	"log"
)

var dbmap *gorp.DbMap

func GetDbMap() *gorp.DbMap {
	return dbmap
}

func InitDb() {
	db, err := sql.Open("mysql", "root:45645655@tcp(localhost:3306)/zipdb")
	checkErr(err, "sql.Open failed")
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(entity.Article{}, "articles").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	//	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
