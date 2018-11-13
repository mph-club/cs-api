package database

import (
	"csportal-server/models"
	"errors"
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
			Order("updated_time ASC").
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
			Order("updated_time ASC").
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

func InsertCarNote(note *models.VehicleNote) error {
	db := connectToDB()

	if err := db.Insert(note); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func InsertUserNote(note *models.UserNote) error {
	db := connectToDB()

	if err := db.Insert(note); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func EditCarNote(note *models.VehicleNote) error {
	db := connectToDB()

	if err := db.Select(note); err != nil {
		log.Println(err)
	}

	err := db.Update(note)
	if err != nil {
		return err
	}

	return nil
}

func EditUserNote(note *models.UserNote) error {
	db := connectToDB()

	if err := db.Select(note); err != nil {
		return err
	}

	err := db.Update(note)
	if err != nil {
		return err
	}

	return nil
}

func GetUserNotes(u *models.User) ([]models.UserNote, error) {
	db := connectToDB()

	var users []models.User
	err := db.Model(&users).
		Column("user.*", "UserNotes").
		Relation("UserNotes", func(q *orm.Query) (*orm.Query, error) {
			return q.Order("created_time DESC"), nil
		}).
		Select()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if len(users[0].UserNotes) == 0 {
		return nil, errors.New("This user has no notes")
	}

	return users[0].UserNotes, nil
}

func GetVehicleNotes(v *models.Vehicle) ([]models.VehicleNote, error) {
	db := connectToDB()

	var vehicles []models.Vehicle

	err := db.Model(&vehicles).
		Column("vehicle.*", "VehicleNotes").
		Relation("VehicleNotes", func(q *orm.Query) (*orm.Query, error) {
			return q.Order("created_time DESC"), nil
		}).
		Where("id = ?", v.ID).
		Select()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if len(vehicles[0].VehicleNotes) == 0 {
		return nil, errors.New("This vehicle has no notes")
	}

	return vehicles[0].VehicleNotes, nil
}
