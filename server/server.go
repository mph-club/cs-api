package server

import "github.com/kataras/iris"

// CreateAndListen exposes the listen and creation of the api
func CreateAndListen() {
	_api := iris.New()

	v1 := _api.Party("api/v1")
	{
		v1.Get("/home", func(ctx iris.Context) {
			ctx.Writef("cs portal home!!!!")
		})

		v1.Get("/getApprovalQueue", cognitoAuth, getApprovalQueue)
		v1.Post("/editCarStatus", editCarStatus)
	}

	_api.Run(iris.Addr(":8081"))
}
