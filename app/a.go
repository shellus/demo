package app

import (
	"fmt"
	"github.com/shellus/demo/app/db"
)

func MainA(){
	fmt.Println("我是a.go的函数,我拿到了DB, 正在ping")
	fmt.Println(db.Session.Ping())
}