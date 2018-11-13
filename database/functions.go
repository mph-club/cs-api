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

func InsertNote(note *models.Note) error {
	db := connectToDB()

	if err := db.Insert(&note); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func EditNote(note *models.Note) error {
	db := connectToDB()

	if err := db.Select(&note); err != nil {
		return err
	}

	err := db.Update(&note)
	if err != nil {
		return err
	}

	return nil
}

func GetUserNotes(u *models.User) ([]models.Note, error) {
	db := connectToDB()

	var user []models.User

	err := db.Model(&user).
		Column("user.*", "Notes").
		Relation("Notes", func(q *orm.Query) (*orm.Query, error) {
			return q.Order("created_time DESC"), nil
		}).
		Where("id = ?", u.ID).
		Select()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user[0].Notes, nil
}

func GetVehicleNotes(v *models.Vehicle) ([]models.Note, error) {
	db := connectToDB()

	var vehicle []models.Vehicle

	err := db.Model(&vehicle).
		Column("vehicle.*", "Notes").
		Relation("Notes", func(q *orm.Query) (*orm.Query, error) {
			return q.Order("created_time DESC"), nil
		}).
		Where("id = ?", v.ID).
		Select()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return vehicle[0].Notes, nil
}
