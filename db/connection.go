package db

import (
	"fmt"
	"log"

	"github.com/kelompok4-loyaltypointagent/backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbConfig := config.LoadDBConfig()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal("Can't connect to database!")
	}

	return db
}
