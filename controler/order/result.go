package order

import (
	"myapp/Model/result"
	"myapp/Model/user"
	configs "myapp/config"
	"myapp/repo/mathematics/kripto"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Result(c echo.Context) error {
	var userName string
	var result result.Result
	var user user.User
	var FoundBuy, FoundSell bool
	id := c.Param("id")
	err := configs.DB.First(&user, id).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, "user tidak ada")
	}
	// siapin database

	FoundBuy, FoundSell = kripto.TransactionExist(id)
	if FoundBuy || FoundSell {
		userName = kripto.FindUserById(id)
		conv, _ := strconv.ParseUint(id, 10, 64)
		result.Conclusion, result.AssetResult = kripto.Findconclusion(kripto.ResultAllSell(id), kripto.ResultAllBuy(id))
		result.UserName = userName
		result.UserId = uint(conv)
		configs.DB.Create(&result)
	} else if !FoundBuy {
		return c.JSON(http.StatusNotFound, "user doest have any transaction")
	}
	return c.JSON(http.StatusOK, result)
}

// var allSell, allBuy int
// configs.DB.Raw("SELECT SUM(intake) FROM sells WHERE user_id = ?", id).Scan(&allSell)
// allSell = kripto.ResultAllSell(id)
// configs.DB.Raw("SELECT SUM(outake) FROM transactions WHERE user_id = ?", id).Scan(&allBuy)
// allBuy = kripto.ResultAllBuy(id)
// consclusions, adding = conclusion(allSell, allBuy)
// configs.DB.Raw("SELECT name FROM users WHERE id = ?", id).Scan(&userName)
