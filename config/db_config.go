package config

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

func InitDbConfig() string {
	filePath := "./config.toml"
	config, err := toml.LoadFile(filePath)
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Get("mysql.user"),
		config.Get("mysql.pass"),
		config.Get("mysql.server"),
		config.Get("mysql.port"),
		config.Get("mysql.schema"))
	return dsn
}
