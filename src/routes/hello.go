package routes

import iris "github.com/kataras/iris/v12"

// HelloWorld say hello with json
func HelloWorld(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "Hello Iris!"})
}
