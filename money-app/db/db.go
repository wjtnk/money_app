/*
db接続処理とマイグレーション
*/
package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB
var err error

const (
	Dialect    = "mysql"
	DBUser     = "root"
	DBPass     = "root"
	DBProtocol = "tcp(mysql:3306)"
	DBName     = "money_app"
)

func Init() {
	connectTemplate := "%s:%s@%s/%s?parseTime=true"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)

	db, err = gorm.Open(Dialect, connect)

	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	Migrate()
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	closeDB := GetDB()
	closeDB.Close()
}
