package user

import (
	"app/config"
	// "io/ioutil"
	"strconv"
)

type User struct {
	tableName struct{}	`pg:"users"`
	Id uint64			`pg:"user_id"		json:"id"`
	Firstname string	`pg:"firstname"		json:"firstName"`
	Lastname string		`pg:"lastname"		json:"lastName"`
	CreatedAt string	`pg:"created_at"	json:"createdAt"`
}

func NewUser (user User) error {
	pg := config.DBConnection()
	_, err := pg.Model(&user).Insert()

	return err
}

func GetById (id string) (User, error) {

	user := User {}
	pg := config.DBConnection()

	err := pg.Model(&user).Where("user_id = ?0", id).Select()
	return user, err
}

func GetUser() ([]User, error) {

	pg := config.DBConnection()

	user := []User {}

	err := pg.Model(&user).Select()

	return user, err
}


func DeleteUser(id string) error {

  user_id, _ := strconv.Atoi(id)

  pg := config.DBConnection() 

  user := []User {}

  _, err := pg.Model(&user).Where("user_id = ?0", user_id).Delete()

  return err
}