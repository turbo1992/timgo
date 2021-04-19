package main

import (
	"tim-go/core"
	"tim-go/global"
	"tim-go/initialize"
)

func main()  {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		initialize.Mysql()
	// case "sqlite":
	//	initialize.Sqlite()  // sqlite需要gcc支持 windows用户需要自行安装gcc 如需使用打开注释即可
	default:
		initialize.Mysql()
	}
	initialize.DBTables()

	defer global.GVA_DB.Close()

	//if err := global.GVA_DB.Raw("select * from casbin_rule").Scan(&cusbin.Lines).Error; err != nil {
	//	fmt.Println("casbin_rule err :", err)
	//}

	//controller.Test()

	core.RunWindowsServer()

	//when timer start, open this code, keep the program alive
	//time.Sleep(1e9*time.Second)
}

