package modelb

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserLogin struct {
	ID              int
	Pwd             string
	Name            string
}

