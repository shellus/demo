package db

import "fmt"
import (
	"gopkg.in/mgo.v2"
)

var Session *mgo.Session

func ConnectDB() {
	fmt.Println("我是db->session.go,我在链接数据库")

	var err error

	Session, err = mgo.Dial("127.0.0.1:40001")
	if err != nil {
		//panic(err)
	}
	// Optional. Switch the session to a monotonic behavior.
	Session.SetMode(mgo.Monotonic, true)
}
func CloseSession(){
	Session.Close()
}

