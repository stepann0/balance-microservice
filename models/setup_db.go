package models

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// user: 'root', password: 'root'
	// host: db:3306, db: bank
	fmt.Println("trying to connect to the DB...")
	dsn := "root:root@tcp(db:3306)/balances?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// пытаемся подключиться к БД пока не получим err == nil
	for err != nil {
		fmt.Println("failed to connect to database. Trying again...")
		time.Sleep(time.Millisecond * 300)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}
	fmt.Println("DB connected!")

	db.Exec("DROP TABLE IF EXISTS accounts, reservation_accounts, services, payments")
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&ReservationAccount{})
	db.AutoMigrate(&Service{})
	db.AutoMigrate(&Payment{})
	DB = db

	CreateServices()
}

func CreateServices() {
	DB.Create(&Service{Name: "покупка"})
	DB.Create(&Service{Name: "консультация"})
	DB.Create(&Service{Name: "ремонт"})
	DB.Create(&Service{Name: "доставка"})
}
