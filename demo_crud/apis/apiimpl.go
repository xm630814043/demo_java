package apis

import (
	"demo_crud/moude"
	"demo_crud/services"
	"fmt"
	"github.com/jinzhu/gorm"
)

func QuerySysConfigbyid(lists *moude.GetType,db *gorm.DB){
	//fmt.Println(lists.ParamType)
	data := services.GetSysConfigByTypebyid(lists,db)
	fmt.Println(data)
}

func QuerySysConfig(db *gorm.DB){
	data := services.GetSysConfigByType(db)
	fmt.Println(len(data),data)
}

/*查询全部*/
//func QuerySysConfig(db *sql.DB) {
//	err :=  services.GetSysConfigByType(db)
//	fmt.Println(err)
//}
//
///*根据param_type查询*/
//func QuerySysConfigById(id int,db *sql.DB) {
//	err :=  services.GetSysConfigByTypeById(id,db)
//	fmt.Println(err)
//}


/*根据id删除*/
func DeleteSysConfig(id int,db *gorm.DB){
	code := services.DeleteSysConfigById(id,db)
	fmt.Println(code)
}

/*添加*/
func SaveSysConfig(lists *moude.SysConfig,db *gorm.DB) {
	//fmt.Println("添加单元测试传输值",lists)
	code := services.AddSysConfig(lists,db)
	fmt.Println(code)
}

//修改
func UpdateSysConfig(id int,params map[string]interface{},db *gorm.DB) {
	lists := &moude.SysConfig{}
	code := services.EditSysConfig(id, lists, params,db)
	fmt.Println(code)
}