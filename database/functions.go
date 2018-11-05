package database

import (
	"csportal-server/models"
	"log"
	"net/url"

	"github.com/go-pg/pg/orm"
)

func GetApprovalQueue(queryParams url.Values) ([]models.Vehicle, error) {
	var vehicleList []models.Vehicle

	db := connectToDB()

	err := db.Model(&vehicleList).
		Apply(orm.Pagination(queryParams)).
		Where("status = ?", "PENDING").
		Select()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return vehicleList, nil
}
