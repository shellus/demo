package app

import (
	"fmt"
	"github.com/shellus/demo/app/db"
)

func MainC(){
	fmt.Printf("我是c.go的函数,我拿到了DB: %d\n", db.DB)
}