package crypto

import (
	"myapp/Model/api"
	"net/http"

	"github.com/labstack/echo/v4"
	gecko "github.com/superoo7/go-gecko/v3"
)

func Cekprice(c echo.Context) error {
	cg := gecko.NewClient(nil)
	id := c.Param("id")
	price, err := cg.SimpleSinglePrice(id, "idr")
	if err != nil {
		return c.String(http.StatusBadRequest, "not found")
	}
	var response = api.Kripto{
		http.StatusOK,
		price.ID,
		price.MarketPrice,
		price.Currency,
	}
	return c.JSON(http.StatusOK, response)
}
