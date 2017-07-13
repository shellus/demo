package app

import (
	"fmt"
	"github.com/shellus/demo/app/db"
)

func MainA(){
	fmt.Printf("我是a.go的函数,我拿到了DB: %d\n", db.DB)
}