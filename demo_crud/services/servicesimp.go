package services

import (
	"demo_crud/moude"
	"github.com/jinzhu/gorm"
)

/*获取*/
func  GetSysConfigByTypebyid(e *moude.GetType,db *gorm.DB) (list []*moude.SysConfig) {
	var lists []*moude.SysConfig
	//fmt.Println(e.ParamType[0])
	err := db.Table("sys_configs").Where("param_type = ?", e.ParamType[0]).Find(&lists).Error
	if err != nil {
		//return errors.New("查询失败！")
	}
	return lists
}

func  GetSysConfigByType(db *gorm.DB) (configs []*moude.SysConfig) {
	var config []*moude.SysConfig
	err := db.Table("sys_configs").Find(&config)
	if err != nil {
		//return errors.New("查询失败！")
	}
	return config
}

/*查询全部*/
//func  GetSysConfigByType(db *sql.DB) int {
//	rows, _ := db.Query("SELECT id,param_key,param_value,param_type FROM sys_configs")
//	for rows.Next() {
//		var id int
//		var param_key string
//		var param_value string
//		var param_type int
//		//Columns列
//		rows.Columns()
//		//read:从数据库中读一条记录。
//		//scan：在数据库中执行范围查询，结果返回一个记录集。
//		_ = rows.Scan(&id, &param_key, &param_value, &param_type)
//		fmt.Println( strconv.Itoa(id) + "\t" +  param_key + "\t" + param_value + "\t" + strconv.Itoa(param_type))
//	}
//	return 200
//}
//
///*根据param_type查询*/
//func  GetSysConfigByTypeById(ids int,db *sql.DB) int {
//	stmts, _ := db.Prepare("SELECT id,param_key,param_value,param_type FROM sys_configs where id=?")
//	rows, _  := stmts.Exec(ids)
//	fmt.Println("根据param_type查询", rows)
//	return 200
//}

/*删除*/
func  DeleteSysConfigById(id int,db *gorm.DB) int {
	err := db.Where("id=?", id).Delete(&moude.SysConfig{}).Error
	if err != nil {
		return 500
	}
	return 200
}

/*添加*/
func  AddSysConfig(sys *moude.SysConfig,db *gorm.DB) int {
	err := db.Create(&sys).Error
	//fmt.Println("sql语句执行返回后的结果",err)
	if err != nil {
		return 500
	}
	return 200
}

/*修改*/
func  EditSysConfig(id int, sys *moude.SysConfig, params map[string]interface{},db *gorm.DB) int {
	//fmt.Println("调用服务层后返回值",params)
	err :=db.Model(&sys).Where("id=?", id).Update(params).Error
	if err != nil {
		return 500
	}
	return 200
}
