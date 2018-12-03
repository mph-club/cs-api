package database

import (
	"csportal-server/models"
	"errors"
	"log"
	"net/url"
	"strings"

	"github.com/go-pg/pg/orm"
)

func UpsertStaff(s models.Staff) error {
	db := connectToDB()

	staff := models.Staff{
		ID: s.ID,
	}

	if err := db.Select(&staff); err != nil {
		log.Println(err.Error())
		log.Println("user does not exist, create")

		staff = staff.Merge(s)
	} else {
		log.Println("user does exist, update")

		s = s.Merge(staff)

		if dbErr := db.Update(&s); dbErr != nil {
			return dbErr
		}
		return err
	}

	if err := db.Insert(&staff); err != nil {
		return err
	}

	return nil
}

func GetStaffList() (int, []models.Staff, error) {
	var staffList []models.Staff

	db := connectToDB()
	count, err := db.Model(&staffList).
		SelectAndCount()

	if err != nil {
		return -1, nil, err
	}

	return count, staffList, nil
}

func GetStaffDetail(s models.Staff) (models.Staff, error) {
	db := connectToDB()

	if err := db.Model(&s).
		WherePK().
		Select(); err != nil {
		return models.Staff{}, err
	}

	return s, nil
}

func DeleteStaff(s models.Staff) error {
	db := connectToDB()

	if _, err := db.Model(&s).
		WherePK().
		Delete(); err != nil {
		return err
	}

	return nil
}

func GetUserList() (int, []models.User, error) {
	var userList []models.User

	db := connectToDB()
	count, err := db.Model(&userList).
		SelectAndCount()

	if err != nil {
		return -1, nil, err
	}

	return count, userList, nil
}

func GetVehicleDetail(v models.Vehicle) (models.Vehicle, error) {
	db := connectToDB()

	if err := db.Model(&v).
		WherePK().
		Select(); err != nil {
		return models.Vehicle{}, err
	}

	return v, nil
}

func GetUserDetail(u models.User) (models.User, error) {
	db := connectToDB()

	if err := db.Model(&u).
		WherePK().
		Select(); err != nil {
		return models.User{}, err
	}

	return u, nil
}

func GetApprovalQueue(queryParams url.Values, status string) (int, []models.Vehicle, error) {
	var vehicleList []models.Vehicle

	db := connectToDB()

	if len(status) == 0 {
		count, err := db.Model(&vehicleList).
			Apply(orm.Pagination(queryParams)).
			Order("updated_time ASC").
			SelectAndCount()
		if err != nil {
			log.Println(err)
			return -1, nil, err
		}

		return count, vehicleList, nil
	} else {
		status = strings.ToUpper(status)

		count, err := db.Model(&vehicleList).
			Apply(orm.Pagination(queryParams)).
			Where("status = ?", status).
			Order("updated_time ASC").
			SelectAndCount()
		if err != nil {
			log.Println(err)
			return -1, nil, err
		}

		return count, vehicleList, nil
	}
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
