package main

import (
	"github.com/hela/config"
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})

	// Listen and serve on http://localhost:8080.
	app.Run(iris.Addr(config.GetConfig().HTTP_PORT))
}
