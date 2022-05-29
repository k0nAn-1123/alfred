package dal

import (
	"alfred/commom/db_init"
	"alfred/model"
)

func Create(user model.User) error {
	db := db_init.InitDB()
	err := db.Create(&user).Error
	if err != nil {
		return nil
	}

	return nil
}
