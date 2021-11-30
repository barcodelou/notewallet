package order_test

import (
	"encoding/json"
	"myapp/Model/api"
	configs "myapp/config"
	"myapp/controler/order"
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	configs.InitDBTest()
	e := echo.New()
	return e
}

// var (
// 	userJSON = `{"userId":2,"coin":"bitcoin","jumlah_beli":10000}`
// )

// func TestBuyCritpo(t *testing.T) {

// 	e := InitEcho()
// 	baserespon := api.BaseResponse{}
// 	var user user.User
// 	var buy transaction.Transaction
// 	t.Run("Success Create user", func(t *testing.T) {
// 		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		rec := httptest.NewRecorder()
// 		err := configs.DB.First(&user, buy.UserId).Error
// 		log.Println(rec)
// 		c := e.NewContext(req, rec)
// 		c.SetPath("/crypto/buy")
// 		if user.ID == 0 {
// 			assert.Equal(t, http.StatusBadRequest, baserespon.Code)
// 		} else if err != nil {
// 			assert.Equal(t, http.StatusInternalServerError, baserespon.Code)
// 		}
// 		if assert.NoError(t, order.Pembelian(c)) {
// 			body := rec.Body.String()
// 			baseResponse := api.BaseResponse{}
// 			if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
// 				assert.Error(t, err, "Failed convert body to object")
// 			}
// 			assert.Equal(t, http.StatusOK, baseResponse.Code)
// 		}
// 	})
// }
// var (
// 	mockDB = map[string]*user{
// 		"jon@labstack.com": &user{"Jon Snow", "jon@labstack.com"},
// 	}

// 	userJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`
// )
// func CreateSeedNews() {
// 	var user user.User
// 	user.Name = "pao"
// 	user.Email = "sanpaulo@gmail.com"
// 	user.Asset = 30000000
// 	configs.DB.Create(&user)
// }
// func CreateSeedNews() {
// 	var user user.User
// 	user.Name = "pao"
// 	user.Email = "sanpaulo@gmail.com"
// 	user.Asset = 30000000
// 	configs.DB.Create(&user)
// }

var (
	userJSON = `{
		"userId":2,
		"jumlah_beli":20000,
		"coin":"bitcoin"
	}`
	// falseJson = `{
	// 	"userId":0,
	// 	"jumlah_beli":20000,
	// 	"coin":"bitcoin"
	// }`
)

func TestBuyCripto(t *testing.T) {

	e := InitEcho()
	// CreateSeedNews()
	t.Run("Success Create transaction", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/transactions/buy")
		if assert.NoError(t, order.Pembelian(c)) {
			body := rec.Body.String()
			baseResponse := api.BaseResponse{}
			if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
				assert.Error(t, err, "Failed convert body to object")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	// t.Run("false input", func(t *testing.T) {
	// 	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(falseJson))
	// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)

	// 	c.SetPath("/transactions/buy")
	// 	if assert.NoError(t, order.Pembelian(c)) {
	// 		body := rec.Body.String()
	// 		baseResponse := api.BaseResponse{}
	// 		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
	// 			assert.Error(t, err, "Failed convert body to object")
	// 		}
	// 		assert.Equal(t, 400, rec.Code)
	// 	}
	// })
}
