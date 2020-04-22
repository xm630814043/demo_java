package storages

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var db *gorm.DB

func DBConn()(db *gorm.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "test?charset=utf8"
	db, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		fmt.Println("连接数据库失败")
	} else {
		fmt.Println("连接数据库成功")
	}
	return db
}

func DBConns()(db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "test?charset=utf8"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		fmt.Println("连接数据库失败")
	} else {
		fmt.Println("连接数据库成功")
	}
	return db
}



