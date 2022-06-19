package dal

import (
	"alfred/commom/db_init"
	"alfred/model"
)

func CreateUser(user model.User) error {
	db := db_init.InitDB()
	err := db.Create(&user).Error
	if err != nil {
		return nil
	}

	return nil
}

func QueryUserById(id int64) (model.User, error) {
	var user model.User
	db := db_init.InitDB()
	err := db.First(&user, id).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
