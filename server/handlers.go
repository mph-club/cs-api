package server

import "github.com/kataras/iris"

func getApprovalQueue(ctx iris.Context) {
	//fetch paged list of cars with status: PENDING
}

func editCarStatus(ctx iris.Context) {
	//change car's status with the request body's enum value
}

func addNotes(ctx iris.Context) {
	//add a note to the users or vehicles notes array
}

func deleteNotes(ctx iris.Context) {
	//deletes a note from the users or vehicles notes array
}
