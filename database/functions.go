package database

import (
	"csportal-server/models"
	"log"
	"net/url"
	"strings"

	"github.com/go-pg/pg/orm"
)

func GetApprovalQueue(queryParams url.Values, status string) ([]models.Vehicle, error) {
	var vehicleList []models.Vehicle

	db := connectToDB()

	if len(status) == 0 {
		err := db.Model(&vehicleList).
			Apply(orm.Pagination(queryParams)).
			Select()
		if err != nil {
			log.Println(err)
			return nil, err
		}
	} else {
		status = strings.ToUpper(status)

		err := db.Model(&vehicleList).
			Apply(orm.Pagination(queryParams)).
			Where("status = ?", status).
			Select()
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return vehicleList, nil
}

func EditCarStatus(vehicle *models.Vehicle) error {
	db := connectToDB()

	_, err := db.Model(vehicle).
		Column("status").
		WherePK().
		Update()

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
