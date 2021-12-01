package configs

import (

	// "library/model/userbuy"

	"fmt"
	"myapp/Model/result"
	"myapp/Model/seller"
	"myapp/Model/transaction"
	"myapp/Model/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// func CnnectDB() {
// 	dsn := "root:V3rr3ll0u.1@tcp(127.0.0.1:3306)/walet?charset=utf8mb4&parseTime=True&loc=Local"
// 	var err error
// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("Database tidak connect")
// 	}
// 	Migration()
// }

func CnnectDB() {
	dsn := "admin:v3rr3ll0u@tcp(kriptowalet.cj8zkyislcf5.us-east-2.rds.amazonaws.com)/kriptowalet?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database tidak connect")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(user.User{}, seller.Sell{}, transaction.Transaction{}, result.Result{})
}

//testing
type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Database string
}

func InitConfigDBTest() ConfigDB {
	var configDB = ConfigDB{
		DB_Username: "admin",
		DB_Password: "v3rr3ll0u",
		DB_Host:     "testing.cj8zkyislcf5.us-east-2.rds.amazonaws.com",
		DB_Port:     "3306",
		DB_Database: "testing",
	}
	return configDB
}

func InitDBTest() {
	configDB := InitConfigDBTest()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		configDB.DB_Username,
		configDB.DB_Password,
		configDB.DB_Host,
		configDB.DB_Port,
		configDB.DB_Database)

	var error error
	DB, error = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB.Create(&user.User{ID: 1})
	if error != nil {
		panic("Database failed connection : " + error.Error())
	}
	MigrationTest()
}

func MigrationTest() {
	DB.Migrator().DropTable(user.User{}, seller.Sell{}, transaction.Transaction{}, result.Result{})
	DB.AutoMigrate(user.User{}, seller.Sell{}, transaction.Transaction{}, result.Result{})
	var users = []user.User{{ID: 3, Name: "nona", Email: "hacker@gmail", Asset: 15000000}}
	var seller = []transaction.Transaction{{ID: 3, UserId: 3, Outake: 1500000,
		PriceBuy: 180000, Qtt: 10, Coin: "bitcoin", RemnantQtt: 10}}
	// , {ID: 2, Name: "lola", Email: "lola@gmail", Asset: 15000000}
	DB.Create(&users)
	DB.Create(&seller)
}
