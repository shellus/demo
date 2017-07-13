package main

import (
	"github.com/shellus/demo/app/db"
	"github.com/shellus/demo/app"
)

func main(){
	// 这里负责链接数据库，初始化DB
	db.ConnectDB()

	// 调用应用入口
	app.Main()
}