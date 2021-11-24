package order

import (
	"fmt"
	"myapp/Model/api"
	"myapp/Model/handlerfalse"
	"myapp/Model/seller"
	"myapp/Model/transaction"
	"myapp/Model/user"
	configs "myapp/config"
	"net/http"

	"github.com/labstack/echo/v4"
	gecko "github.com/superoo7/go-gecko/v3"
)

func Penjualan(c echo.Context) error {

	var user user.User
	var countqtt, hasil float64
	var seller seller.Sell
	var total int
	c.Bind(&seller)
	cg := gecko.NewClient(nil)
	//for showing a price crypto id
	price, _ := cg.SimpleSinglePrice(seller.Coin, "idr")
	err := configs.DB.First(&user, seller.UserId).Error
	//if user not found
	if err != nil {
		return c.JSON(http.StatusBadRequest, "user tidak ada")
	}
	//sum for remnant qtt(remnant qtt is value that show a rest api counting user asset)
	configs.DB.Raw("SELECT SUM(remnant_qtt) FROM transactions WHERE coin = ? AND user_id = ?", seller.Coin, seller.UserId).Scan(&hasil)
	//counting
	configs.DB.Raw("SELECT COUNT(qtt) FROM transactions WHERE coin = ? AND user_id = ?", seller.Coin, seller.UserId).Scan(&countqtt)
	//handler when found some error input/error condition
	falseInput := handlerfalse.HandlerSell(hasil, seller.Percentage, countqtt)
	if len(falseInput) != 0 {
		return c.JSON(http.StatusBadRequest, falseInput)
	}
	//mathematic formula for getting a variable
	seller.Qtt = seller.Percentage * hasil
	seller.Intake = int(hasil * float64(price.MarketPrice) * seller.Percentage)
	seller.PriceSell = int(price.MarketPrice)
	updatedValue := (hasil - (seller.Qtt)) / countqtt
	//---------------------------------------
	//when we sell some asset so remnant qtt will updated with update value
	sementara := configs.DB.Model(transaction.Transaction{}).Where("coin = ? AND user_id = ?", seller.Coin, seller.UserId).Updates(map[string]interface{}{
		"remnant_qtt": updatedValue,
	})
	configs.DB.Create(&seller)
	//sum for finding a asset money that user got from selling their crypto
	configs.DB.Raw("SELECT SUM(intake) FROM sells WHERE position = ? AND user_id = ?", "aktif", seller.UserId).Scan(&total)
	//updated asset user
	checker := configs.DB.Model(&user).Where("id= ?", seller.UserId).Update("asset", user.Asset+total).Error
	if checker == nil {
		//for make selling position inaktif
		temporary := configs.DB.Model(&seller).Where("position= ? AND user_id = ?", "aktif", seller.UserId).Updates(map[string]interface{}{
			"position": "deaktif",
		})
		fmt.Println("update position ", temporary.Error)
	}
	fmt.Println("cek error", sementara.Error)
	var response = api.BaseResponse{
		Code:    http.StatusOK,
		Message: "transaction done",
		Data: map[string]interface{}{
			"coin": seller.Coin,
			"qtt":  seller.Qtt,
		},
	}
	return c.JSON(http.StatusOK, response)
}
