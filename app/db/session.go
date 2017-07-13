package db

import "fmt"

var DB int

func ConnectDB(){
	fmt.Println("我是db->session.go,我在链接数据库")
	DB = 1
}