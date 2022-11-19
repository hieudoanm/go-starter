package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfigs struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Name     string `json:"name"`
	Mode     string `json:"mode"`
	TimeZone string `json:"timeZone"`
}

var postgresDatabase *gorm.DB

func Open(configs DatabaseConfigs) *gorm.DB {
	if postgresDatabase != nil {
		return postgresDatabase
	}
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		configs.Host,
		configs.Port,
		configs.User,
		configs.Pass,
		configs.Name,
		configs.Mode,
		configs.TimeZone,
	)
	database, databaseError := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if databaseError != nil {
		panic(databaseError)
	}
	postgresDatabase = database
	return postgresDatabase
}
