package main

import (
	"demo_crud/apis"
	"demo_crud/moude"
	"demo_crud/pkgs/enums"
	"demo_crud/storages"
)


func main() {
	db := storages.DBConn()
	defer db.Close()

	//调取添加
	lists := &moude.SysConfig{ID:5, ParamKey: "赵六", ParamValue: "赵六", ParamType: 6}
	apis.SaveSysConfig(lists,db)

	//调取删除
	//apis.DeleteSysConfig(7,db)

	//调取修改
	//var ms map[string]interface{}
	//ms = make(map[string]interface{})
	//ms["ParamKey"]="赵六"
	//ms["ParamValue"]="赵六"
	//ms["ParamType"]=6
	//apis.UpdateSysConfig(5,ms,db)

	//调取根据id查询
	gettypes := &moude.GetType{}
	gettypes.ParamType = make([]enums.Config,1)
	apis.QuerySysConfigbyid(gettypes,db)

	//调取查询全部
	apis.QuerySysConfig(db)


	//dbs := storages.DBConns()
	//defer dbs.Close()
	//apis.QuerySysConfig(dbs)

	//apis.QuerySysConfigById(0,dbs)




















	//context引用web.RequestContexts{}的地址，也可以指其指针
	//context := &webs.RequestContexts{Host:"1",Token:"1"}
	//NewSysConfigService传参不接受context指针类型，接受地址类型
	//srvices := servires.NewSysConfigService(context)

	////调取查询
	//gettypes := &moudels.Types{}
	//gettypes.ParamType = make([]enums.Config,0)
	//
	////GetSysConfigByType方法传参变量属性为切片
	//identifier,err := srvices.GetSysConfigByType(gettypes)
	//fmt.Println(identifier,err)

	//调取添加
	//lists := &moudels.Configs{ID: 4, ParamKey: "赵六", ParamValue: "赵六", ParamType: 5}
	//srvices.AddSysConfig(lists)

	////调取修改
	//sysconfig := &moudels.Configs{}
	//var ms map[string]interface{}
	//ms = make(map[string]interface{})
	//ms["ParamKey"]=sysconfig.ParamKey
	//ms["ParamValue"]=sysconfig.ParamValue
	//ms["ParamType"]=sysconfig.ParamType
	//srvices.EditSysConfig(0, sysconfig,ms)

}

