package server

import (
	"csportal-server/database"
	"csportal-server/models"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func getUserList(ctx echo.Context) error {
	count, data, err := database.GetUserList()

	if err != nil {
		return ctx.JSON(
			response(false, http.StatusBadRequest, map[string]interface{}{"db_error": err.Error()}))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{
				"Users": data,
				"Count": count,
			},
		))
}

func getApprovalQueue(ctx echo.Context) error {
	urlQuery := ctx.Request().URL.Query()
	status := ctx.QueryParam("status")

	//fetch paged list of cars with status: PENDING
	count, data, err := database.GetApprovalQueue(urlQuery, status)

	if err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"db_error": err.Error()}))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{
				"Vehicles": data,
				"Count":    count,
			},
		))
}

func editCarStatus(ctx echo.Context) error {
	//change car's status with the request body's enum value
	var v models.Vehicle

	if err := ctx.Bind(&v); err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"database_error": err.Error()}))
	}

	err := database.EditCarStatus(&v)
	if err != nil {
		return ctx.JSON(
			response(false, http.StatusBadRequest, map[string]interface{}{"db_error": err.Error()}))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{
				"result": "vehicle status was updated",
			},
		))
}

func getVehicleDetail(ctx echo.Context) error {
	var v models.Vehicle
	v.ID = ctx.Param("id")

	detail, err := database.GetVehicleDetail(v)
	if err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"db_error": err.Error()}))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{
				"Vehicle": detail,
			},
		))
}

func getUserDetail(ctx echo.Context) error {
	var u models.User
	u.ID = ctx.Param("id")

	detail, err := database.GetUserDetail(u)
	if err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"db_error": err.Error()}))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{
				"User": detail,
			},
		))
}

func addCarNote(ctx echo.Context) error {
	//add a note to the users or vehicles notes array
	var n models.VehicleNote

	if err := ctx.Bind(&n); err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"server_error": err.Error()}))
	}

	err := database.InsertCarNote(&n)
	if err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"db_error": err.Error()}))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{
				"result": "note was inserted",
			},
		))
}

func addUserNote(ctx echo.Context) error {
	//add a note to the users or vehicles notes array
	var n models.UserNote

	if err := ctx.Bind(&n); err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"server_error": err.Error()}))
	}

	err := database.InsertUserNote(&n)

	if err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"db_error": err.Error()}))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{
				"result": "note was inserted",
			},
		))
}

func getNotesForCar(ctx echo.Context) error {
	var v models.Vehicle

	v.ID = ctx.QueryParam("id")

	list, err := database.GetVehicleNotes(&v)
	if err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"database_error": err.Error()}))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{
				"notes": list,
			},
		))
}

func getNotesForUser(ctx echo.Context) error {
	var u models.User

	u.ID = ctx.QueryParam("id")

	list, err := database.GetUserNotes(&u)
	if err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"database_error": err.Error()}))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{
				"notes": list,
			},
		))
}

func deleteNotes(ctx echo.Context) error {
	//deletes a note from the users or vehicles notes array
	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{
				"result": "delete successful",
			},
		))
}

func upsertStaff(ctx echo.Context) error {
	var s models.Staff

	if err := ctx.Bind(&s); err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"json_bind_error": err.Error()}))
	}

	log.Println(s)

	if err := database.UpsertStaff(s); err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"database_error": err.Error()}))
	}

	resultString := fmt.Sprintf("user was successfully updated")

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{"result": resultString},
		))
}

func getStaff(ctx echo.Context) error {
	count, data, err := database.GetStaffList()

	if err != nil {
		return ctx.JSON(
			response(false, http.StatusBadRequest, map[string]interface{}{"db_error": err.Error()}))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{
				"Staff": data,
				"Count": count,
			},
		))
}

func getStaffDetail(ctx echo.Context) error {
	var s models.Staff
	s.ID = ctx.Param("id")

	detail, err := database.GetStaffDetail(s)
	if err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"db_error": err.Error()}))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{
				"Staff": detail,
			},
		))
}

func deleteStaff(ctx echo.Context) error {
	var s models.Staff
	s.ID = ctx.Param("id")

	if err := database.DeleteStaff(s); err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"db_error": err.Error()}))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{
				"result": "delete successful",
			},
		))
}
