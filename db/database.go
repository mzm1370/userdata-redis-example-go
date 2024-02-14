package db

import (
	"userdata-redis-example-go/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Initialize(cfg *config.Config) (*gorm.DB, error) {
	var err error

	dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBName + "?parseTime=true"

	dbInstance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db = dbInstance

	// if err := db.AutoMigrate(&models.User{}); err != nil {
	// 	return nil, err
	// }

	return db, nil

}

func GetDB() *gorm.DB {
	return db
}

func Close() error {
	dbSQL, err := db.DB()
	if err != nil {
		return err
	}

	return dbSQL.Close()
}
