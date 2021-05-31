package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"money-app/db"
	"money-app/server"
)

func main() {
	db.Init()
	server.Init()
	db.CloseDB()
}
