package app

import (
	"fmt"
	"github.com/shellus/demo/app/db"
)

func MainB(){
	fmt.Printf("我是b.go的函数,我拿到了DB: %d\n", db.DB)
}