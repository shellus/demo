package app

import (
	"fmt"
	"github.com/shellus/demo/app/db"
)

func MainB(){
	fmt.Println("我是b.go的函数,我拿到了DB, 正在ping")
	fmt.Println(db.Session.Ping())
}