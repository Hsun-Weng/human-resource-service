package util

import (
	"fmt"
	"github.com/Hsun-Weng/human-resource-service/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewDb() *gorm.DB {
	dbHost := config.GetEnv(config.DatabaseHost, "localhost")
	dbPort := config.GetEnv(config.DatabasePort, "3306")
	dbUser := config.GetEnv(config.DatabaseUser, "root")
	dbPassword := config.GetEnv(config.DatabasePassword, "human_resource")
	dbName := config.GetEnv(config.DatabaseName, "human_resource")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", dbUser, dbPassword, dbHost, dbPort, dbName)
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	return db
}
