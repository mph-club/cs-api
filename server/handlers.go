package server

import (
	"csportal-server/database"
	"csportal-server/models"
	"net/http"

	"github.com/labstack/echo"
)

func getApprovalQueue(ctx echo.Context) error {
	urlQuery := ctx.Request().URL.Query()
	status := ctx.FormValue("status")

	//fetch paged list of cars with status: PENDING
	data, err := database.GetApprovalQueue(urlQuery, status)

	if err != nil {
		return ctx.JSON(
			makeErrorResponse(
				http.StatusBadRequest,
				map[string]interface{}{"db_error": err.Error()}))
	}

	return ctx.JSON(makeOKResponse(map[string]interface{}{"Vehicles": data}))
}

func editCarStatus(ctx echo.Context) error {
	//change car's status with the request body's enum value
	var v models.Vehicle

	if err := ctx.Bind(&v); err != nil {
		return ctx.JSON(
			makeErrorResponse(
				http.StatusBadRequest,
				map[string]interface{}{"database_error": err.Error()}))
	}

	err := database.EditCarStatus(&v)

	if err != nil {
		return ctx.JSON(
			makeErrorResponse(
				http.StatusBadRequest,
				map[string]interface{}{"db_error": err.Error()}))
	}

	return ctx.JSON(makeOKResponse(map[string]interface{}{"result": "vehicle status was updated"}))
}

func addNote(ctx echo.Context) error {
	//add a note to the users or vehicles notes array
	var n models.Note

	if err := ctx.Bind(&n); err != nil {
		return ctx.JSON(
			makeErrorResponse(
				http.StatusBadRequest,
				map[string]interface{}{"server_error": err.Error()}))
	}

	err := database.InsertNote(&n)

	if err != nil {
		return ctx.JSON(
			makeErrorResponse(
				http.StatusBadRequest,
				map[string]interface{}{"db_error": err.Error()}))
	}

	return ctx.JSON(makeOKResponse(map[string]interface{}{"result": "note was inserted"}))
}

func deleteNotes(ctx echo.Context) {
	//deletes a note from the users or vehicles notes array
}
