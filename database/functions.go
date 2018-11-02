package database

import (
	"log"
	"mphclub-rest-server/models"
)

func GetApprovalQueue() ([]models.Vehicle, error) {
	var vehicleList []models.Vehicle

	db := connectToDB()

	err := db.Model(&vehicleList).Select()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return vehicleList, nil
}
