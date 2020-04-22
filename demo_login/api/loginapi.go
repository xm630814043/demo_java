package api

import (
	"demo_login/modelb"
	"errors"
	"github.com/jinzhu/gorm"
)

func getLink() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	return db
}

func SaveUser(user *modelb.UserLogin) error {
	db := getLink()
	//fmt.Println("注册用户",user)
	err := db.Table("user_login").Create(user).Error
	if err != nil {
		return errors.New("用户名或密码错误！")
	}
	defer db.Close()
	return err
}

func ValidateUser(user *modelb.UserLogin) error {
	//fmt.Println("接受到的传参Name",user.Name,"接受到的传参Pwd",user.Pwd)
	db := getLink()
	err :=db.Table("user_login").Where("name=? and pwd=?", user.Name, user.Pwd).Find(user).Error
	if err != nil {
		return errors.New("用户名或密码错误！")
	}
	defer db.Close()
	return err
}

func QuerydateUser(pagesize,pageno int) (userlo []*modelb.UserLogin) {
	//fmt.Println("接受到的传参Name",user.Name,"接受到的传参Pwd",user.Pwd)
	db := getLink()
	var users []*modelb.UserLogin
	err :=db.Table("user_login").Find(&users)
	if err != nil {
		//return errors.New("查询失败！")
	}
	defer db.Close()
	return users
}


