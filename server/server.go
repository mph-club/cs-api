package server

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

// CreateAndListen exposes the listen and creation of the api
func CreateAndListen() {
	_api := iris.New()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})

	v1 := _api.Party("api/v1", cognitoAuth, crs).AllowMethods(iris.MethodOptions)
	{
		v1.Get("/home", func(ctx iris.Context) {
			ctx.Writef("cs portal home!!!!")
		})

		v1.Get("/getAll", getApprovalQueue)
		v1.Post("/editCarStatus", editCarStatus)
	}

	_api.Run(iris.Addr(":8081"))
}
