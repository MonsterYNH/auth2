package database

import (
	"github.com/MonsterYNH/athena/config"
	"github.com/MonsterYNH/athena/database"
	"gorm.io/gorm"
)

var db *database.Database

func init() {
	var err error

	conf, err := config.GetConfig()
	if err != nil {
		panic("[ERROR] init config failed, error: " + err.Error())
	}
	db, err = database.NewDatabaseWithOption(
		&database.PostgresDataBase{},
		func(dc *config.DatabaseConfig) error {
			dc.Host = conf.DatabaseConfig.Host
			dc.Port = conf.DatabaseConfig.Port
			dc.User = conf.DatabaseConfig.User
			dc.Password = conf.DatabaseConfig.Password
			dc.Name = conf.DatabaseConfig.Name
			dc.SSLMode = conf.DatabaseConfig.SSLMode
			dc.TimeZone = conf.DatabaseConfig.TimeZone
			return nil
		},
	)

	if err != nil {
		panic("init db failed")
	}
}

func GetDatabase() *gorm.DB {
	return db.GetConnect()
}
