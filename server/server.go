package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// CreateAndListen exposes the listen and creation of the api
func CreateAndListen() {
	_api := echo.New()

	_api.Use(middleware.Logger())

	allowedMethods := append(middleware.DefaultCORSConfig.AllowMethods, http.MethodOptions)

	_api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     allowedMethods,
		AllowHeaders:     []string{"Authorization", "Content-Type", "Origin", "User-Agent", "Host"},
		ExposeHeaders:    []string{"Authorization", "Content-Type", "Origin", "User-Agent", "Host"},
		AllowCredentials: true,
	}))

	v1 := _api.Group("api/v1", cognitoAuth)
	v1.GET("/home", func(ctx echo.Context) error {
		return ctx.String(200, "cs portal home!!!!")
	})

	v1.GET("/vehicles", getApprovalQueue)
	v1.POST("/editCarStatus", editCarStatus)
	v1.POST("/addCarNote", addCarNote)
	v1.POST("/addUserNote", addUserNote)
	v1.GET("/getCarNotes", getNotesForCar)
	v1.GET("/getUserNotes", getNotesForUser)
	v1.GET("/users", getUserList)
	v1.GET("/vehicles/:id", getVehicleDetail)
	v1.GET("/users/:id", getUserDetail)

	// *** staff endpoints ***
	v1.POST("/staff", upsertStaff)
	v1.GET("/staff", getStaff)
	v1.DELETE("/staff/:id", deleteStaff)

	_api.Start(":8081")
}
