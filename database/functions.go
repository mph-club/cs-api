package database

import (
	"csportal-server/models"
	"log"
)

func GetApprovalQueue() ([]models.Vehicle, error) {
	var vehicleList []models.Vehicle

	db := connectToDB()

	err := db.Model(&vehicleList).
		Where("status = ?", "PENDING").
		Select()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return vehicleList, nil
}
