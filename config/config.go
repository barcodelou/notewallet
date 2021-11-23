package configs

import (

	// "library/model/userbuy"

	"myapp/Model/result"
	"myapp/Model/seller"
	"myapp/Model/transaction"
	"myapp/Model/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CnnectDB() {
	dsn := "root:V3rr3ll0u.1@tcp(127.0.0.1:3306)/walet?charset=utf8mb4&parseTime=True&loc=Local"
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
