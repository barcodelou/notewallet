package order_test

import (
	"encoding/json"
	"myapp/Model/api"
	"myapp/controler/order"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	// sellJSON = `{"userId":1,"coin":"dogecoin","jumlah_jual":1}`
	sellJSON = `{
		"userId":3,
		"jumlah_jual":0.2,
		"coin":"bitcoin"
	}`
)

func TestSellCripto(t *testing.T) {
	e := InitEcho()
	// CreateSeedUser()
	// var seller seller.Sell

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(sellJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/transactions/sell")

	if assert.NoError(t, order.Penjualan(c)) {
		body := rec.Body.String()
		baseResponse := api.BaseResponse{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}
		assert.Equal(t, http.StatusOK, rec.Code)
	}

}
