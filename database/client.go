package database

import (
	"crud/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func Connect(connectionString string) error {
	var err error
	Db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Println("error in database.Connect")
		return err
	}
	log.Println("Connection DB is Successful")
	errMig := Db.AutoMigrate(&entity.Person{})
	if errMig != nil {
		return errMig
	}
	return nil

}

func Migrate(table *entity.Person) {
	err := Db.AutoMigrate(&table)
	if err != nil {
		return
	}
	log.Println("Table migrated")
}
