package model

import (
	"alfred/constant/user_constant"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type User struct {
	ID         int64                    `gorm:"column:id"`
	Name       string                   `gorm:"column:name"`
	Age        int64                    `gorm:"column:age"`
	Gender     user_constant.Gender     `gorm:"column:gender"`
	Permission user_constant.Permission `gorm:"column:permission"`
	Extra      UserExtra                `gorm:"Type:json;column:extra;embedded"`
}

// UserExtra
// gorm 自定义数据类型
// https://gorm.io/zh_CN/docs/data_types.html
type UserExtra struct {
	PhoneNumber []string `json:"phone_number"`
	Address     string   `json:"address"`
}

func (e *UserExtra) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := UserExtra{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return err
	}
	return nil
}

func (e UserExtra) Value() (driver.Value, error) {
	if bytes, err := json.Marshal(e); err != nil {
		return nil, err
	} else {
		return string(bytes), nil
	}
}
