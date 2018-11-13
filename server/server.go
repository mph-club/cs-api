package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// CreateAndListen exposes the listen and creation of the api
func CreateAndListen() {
	_api := echo.New()

	_api.Use(middleware.Logger())

	v1 := _api.Group("api/v1", cognitoAuth)
	v1.GET("/home", func(ctx echo.Context) error {
		return ctx.String(200, "cs portal home!!!!")
	})

	v1.GET("/getAll", getApprovalQueue)
	v1.POST("/editCarStatus", editCarStatus)
	v1.POST("/addNote", addNote)
	v1.GET("/getCarNotes", getNotesForCar)
	v1.GET("/getUserNotes", getNotesForUser)

	_api.Start(":8081")
}
