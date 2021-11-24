package order

import (
	"fmt"
	"myapp/Model/api"
	"myapp/Model/seller"
	"myapp/Model/transaction"
	"myapp/Model/user"
	configs "myapp/config"
	"net/http"

	"github.com/labstack/echo/v4"
	gecko "github.com/superoo7/go-gecko/v3"
	// "gorm.io/gorm"
)

func Penjualan(c echo.Context) error {

	var user user.User
	// var result *gorm.DB
	var countqtt float64
	var seller seller.Sell
	var hasil float64
	var counting int
	var total int
	falseInput := []string{}
	c.Bind(&seller)
	cg := gecko.NewClient(nil)
	price, _ := cg.SimpleSinglePrice(seller.Coin, "idr")
	err := configs.DB.First(&user, seller.UserId).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, "user tidak ada")
	}
	configs.DB.Raw("SELECT SUM(remnant_qtt) FROM transactions WHERE coin = ? AND user_id = ?", seller.Coin, seller.UserId).Scan(&hasil)
	configs.DB.Raw("SELECT COUNT(qtt) FROM transactions WHERE coin = ?", seller.Coin).Scan(&counting)
	configs.DB.Raw("SELECT COUNT(qtt) FROM transactions WHERE coin = ? AND user_id = ?", seller.Coin, seller.UserId).Scan(&countqtt)
	if seller.Percentage > 1 {
		falseInput = append(falseInput, "to much percentage must below 1")
	}
	if counting == 0 {
		falseInput = append(falseInput, "there is no history for purchase that coin")
	}
	if len(falseInput) != 0 {
		return c.JSON(http.StatusBadRequest, falseInput)
	}
	seller.Qtt = seller.Percentage * hasil
	seller.Intake = int(hasil * float64(price.MarketPrice) * seller.Percentage)
	seller.PriceSell = int(price.MarketPrice)
	temp := (hasil - (seller.Qtt)) / countqtt
	sementara := configs.DB.Model(transaction.Transaction{}).Where("coin = ? AND user_id = ?", seller.Coin, seller.UserId).Updates(map[string]interface{}{
		"remnant_qtt": temp,
	})
	configs.DB.Create(&seller)

	configs.DB.Raw("SELECT SUM(intake) FROM sells WHERE position = ? AND user_id = ?", "aktif", seller.UserId).Scan(&total)
	configs.DB.Model(&user).Where("id= ?", seller.UserId).Update("asset", user.Asset+total)
	rekam := configs.DB.Model(&user).Where("id= ?", seller.UserId).Update("asset", user.Asset+total).Error
	if rekam == nil {
		fmt.Println("pertambahan", rekam)
	}

	cekcek := configs.DB.Model(&user).Where("id= ?", seller.UserId).Update("asset", user.Asset+total).Error
	if cekcek == nil {
		temporary := configs.DB.Model(&seller).Where("position= ? AND user_id = ?", "aktif", seller.UserId).Updates(map[string]interface{}{
			"position": "deaktif",
		})
		fmt.Println("update position ", temporary.Error)
	}
	fmt.Println("cek error", sementara.Error)
	// result=configs.DB.Create(&seller)
	// if result != nil || sementara.Error != nil {
	// 	return c.JSON(http.StatusBadRequest, api.BaseResponse{
	// 		http.StatusInternalServerError,
	// 		result.Error.Error(),
	// 		nil,
	// 	})
	// }
	var response = api.BaseResponse{
		http.StatusOK,
		"transaction done",
		map[string]interface{}{
			"coin": seller.Coin,
			"qtt":  seller.Qtt,
		},
	}

	return c.JSON(http.StatusOK, response)
}
