package main

import "github.com/kataras/iris"

func main() {
	app := iris.Default()
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello iris web framework."})
	})
	app.Run(iris.Addr(":8080"))
}
