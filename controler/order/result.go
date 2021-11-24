package order

import (
	"myapp/Model/result"
	"myapp/Model/user"
	configs "myapp/config"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Result(c echo.Context) error {
	var allSell, allBuy, adding int
	var consclusions, userName string
	var result result.Result
	var user user.User
	id := c.Param("id")
	err := configs.DB.First(&user, id).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, "user tidak ada")
	}
	configs.DB.Raw("SELECT SUM(intake) FROM sells WHERE user_id = ?", id).Scan(&allSell)
	configs.DB.Raw("SELECT SUM(outake) FROM transactions WHERE user_id = ?", id).Scan(&allBuy)
	adding = allSell - allBuy
	if adding < 0 {
		consclusions = "loss"
	} else if adding == 0 {
		consclusions = "break event point"
	} else {
		consclusions = "profit"
	}
	configs.DB.Raw("SELECT name FROM users WHERE id = ?", id).Scan(&userName)
	conv, _ := strconv.ParseUint(id, 10, 64)
	result.Conclusion = consclusions
	result.UserName = userName
	result.UserId = uint(conv)
	result.AssetResult = adding
	configs.DB.Create(&result)
	return c.JSON(http.StatusOK, result)
}
