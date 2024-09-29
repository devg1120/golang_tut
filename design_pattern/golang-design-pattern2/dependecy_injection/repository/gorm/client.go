package gorm

import (
	"log"

	"github.com/cocoide/golang-design-pattern/dependecy_injection/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBClient(config *entity.DBConfig) *gorm.DB {
	db, err := gorm.Open(
		mysql.Open(config.DSN),
	)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
