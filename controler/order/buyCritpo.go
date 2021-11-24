package order

import (
	"myapp/Model/api"
	"myapp/Model/transaction"
	"myapp/Model/user"
	configs "myapp/config"
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
	if err != nil {
		return c.JSON(http.StatusBadRequest, "user tidak ada")
	}

	if buy.Outake < user.Asset {
		cg := gecko.NewClient(nil)
		price, _ := cg.SimpleSinglePrice(buy.Coin, "idr")

		// buy.Qtt = float64(buy.Outake) / float64(price.MarketPrice)
		buy.Qtt = mathematics.FindQtt(float64(buy.Outake), float64(price.MarketPrice))
		// buy.RemnantQtt = float64(buy.Outake) / float64(price.MarketPrice)
		buy.RemnantQtt = mathematics.FindQtt(float64(buy.Outake), float64(price.MarketPrice))
		buy.PriceBuy = int(price.MarketPrice)
		result = configs.DB.Create(&buy)
		configs.DB.Model(&user).Where("id= ?", buy.UserId).Update("asset", user.Asset-buy.Outake)
		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, api.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: result.Error.Error(),
				Data:    nil,
			})
		}
	}
	var response = api.BaseResponse{
		Code:    http.StatusOK,
		Message: "transaction done",
		Data: map[string]interface{}{
			"coin": buy.Coin,
			"qtt":  buy.Qtt,
		},
	}
	return c.JSON(http.StatusOK, response)
}
