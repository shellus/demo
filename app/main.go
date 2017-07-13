package app

import (
	"fmt"
	"github.com/shellus/demo/app/db"
)

func Main()  {
	fmt.Printf("我是main.go的Main函数,我拿到了DB: %d\n", db.DB)
	MainA()
	MainB()
	MainC()
}