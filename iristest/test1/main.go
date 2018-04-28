package main

import (
	"github.com/kataras/iris"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
	"fmt"
)


type User struct {
	ID 			int64  		// xorm 设置自动增长
	Version 	string  `xorm:"varchar(200)"`
	Salt 		string
	Username	string
	Password	string  `xorm:"varchar(200)"`
	Languages  	string  `xorm:"varchar(200)"`
	CreatedAt	time.Time	`xorm:"created"`
	UpdatedAt	time.Time 	`xorm:"updated"`
}

func main() {
	app := iris.New()

	orm, err := xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		fmt.Println(err.Error())
	}

	//RegisterOnInterrupt registers a global function to call when CTRL+C/CMD+C pressed or a unix kill command received.
	iris.RegisterOnInterrupt(func() {
		orm.Close()
	})

	// Sync2 synchronize structs to database tables
	err = orm.Sync2(new(User))
	if err != nil {
		fmt.Println(err.Error())
	}

	app.Get("/register", func(ctx iris.Context) {
		user := &User{
			Username:"test",
			Salt:"test",
			Password: "test",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		orm.Insert(user)
		// 返回信息
		ctx.Writef("user inserted: %#v", user)
	})

	app.Get("/get", func(ctx iris.Context) {
		user := User{ID:1}
		if ok, _ := orm.Get(&user); ok {
			ctx.Writef("user found: %#v", user)
		}
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

}
