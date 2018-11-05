package server

import (
	"csportal-server/database"
	"log"

	"github.com/kataras/iris"
)

func getApprovalQueue(ctx iris.Context) {
	urlQuery := ctx.Request().URL.Query()

	//fetch paged list of cars with status: PENDING
	data, err := database.GetApprovalQueue(urlQuery)

	if err != nil {
		log.Println(err)
		ctx.JSON(generateJSONResponse(false, iris.Map{"db_error": err.Error()}))
		return
	}

	ctx.JSON(generateJSONResponse(true, iris.Map{"Vehicles": data}))
	return
}

func editCarStatus(ctx iris.Context) {
	//change car's status with the request body's enum value
	vehicleID := ctx.PostValue("id")
	statusChange := ctx.PostValue("edit_status")

	err := database.EditCarStatus(vehicleID, statusChange)

	if err != nil {
		log.Println(err)
		ctx.JSON(generateJSONResponse(false, iris.Map{"db_error": err.Error()}))
		return
	}

	ctx.JSON(generateJSONResponse(true, iris.Map{"result": "vehicle status was updated"}))
}

func addNotes(ctx iris.Context) {
	//add a note to the users or vehicles notes array
}

func deleteNotes(ctx iris.Context) {
	//deletes a note from the users or vehicles notes array
}
