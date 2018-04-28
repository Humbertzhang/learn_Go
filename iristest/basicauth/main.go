package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/basicauth"
	"time"
)

func App() *iris.Application {
	app := iris.New()

	authconfig := basicauth.Config{
		Users: map[string]string {"username":"password"},
		Realm: "Authorization Required",
		Expires: time.Duration(30) * time.Hour,
	}

	authentication := basicauth.New(authconfig)

	app.Get("/", func(ctx iris.Context) {
		ctx.Redirect("/admin")
	})

	needAuth := app.Party("/admin", authentication)
	{
		needAuth.Get("/", h)
		needAuth.Get("/profile", h)
		needAuth.Get("/settings", h)
	}
	return app
}

func h(ctx iris.Context) {
	username, password, _ := ctx.Request().BasicAuth()
	ctx.Writef("%s %s:%s", ctx.Path(), username, password)
}

func main() {
	app := App()
	app.Run(iris.Addr(":8080"))
}
