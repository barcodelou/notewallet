package order

import (
	"log"
	"myapp/Model/api"
	"myapp/Model/transaction"
	"myapp/Model/user"
	configs "myapp/config"
	handlerespon "myapp/controler/order/handleRespon"
	"myapp/repo/mathematics"
	"net/http"

	"github.com/labstack/echo/v4"
	gecko "github.com/superoo7/go-gecko/v3"
	"gorm.io/gorm"
)

func Pembelian(c echo.Context) error {
	var user user.User
	var result *gorm.DB
	var buy transaction.Transaction
	c.Bind(&buy)
	err := configs.DB.First(&user, buy.UserId).Error
	log.Println(user)
	if user.ID == 0 {
		return c.JSON(http.StatusBadRequest, "user tidak ada")
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, "database error")
	}
	if buy.Outake < user.Asset {
		cg := gecko.NewClient(nil)
		price, _ := cg.SimpleSinglePrice(buy.Coin, "idr")
		buy.Qtt = mathematics.FindQtt(float64(buy.Outake), float64(price.MarketPrice))
		buy.RemnantQtt = mathematics.FindQtt(float64(buy.Outake), float64(price.MarketPrice))
		buy.PriceBuy = int(price.MarketPrice)
		result = configs.DB.Create(&buy)
		configs.DB.Model(&user).Where("id= ?", buy.UserId).Update("asset", user.Asset-buy.Outake)
		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, handlerespon.ErrorBuy(api.BaseResponse{}, result.Error.Error()))
		}
	}
	return c.JSON(http.StatusOK, handlerespon.SucsessBuy(api.BaseResponse{}, buy))
}
