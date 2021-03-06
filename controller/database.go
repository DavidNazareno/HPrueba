package controller

import (
	"log"

	"github.com/DavidNazareno/h_prueba/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ORM struct {
	DB *gorm.DB
}

var Database *gorm.DB

var pathFile = config.DB

//GetDatabase funcion para cargar la base de datos
func LoadDatabase() {

	db, err := gorm.Open(sqlite.Open(pathFile), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), //logger.Silent > disable Log, logger.info show query to database
	})

	db.Debug()
	if err != nil {
		log.Panic(err.Error())
	}
	Database = db

}
