package server

import "github.com/kataras/iris"

// CreateAndListen exposes the listen and creation of the api
func CreateAndListen() {
	_api := iris.New()

	v1 := _api.Party("api/v1", cognitoAuth)
	{
		v1.Get("/", func(ctx iris.Context) {
			ctx.Writef("cs portal home!!!!")
		})

		v1.Get("/getCars", getCars)
	}

	_api.Run(iris.Addr(":8081"))
}
