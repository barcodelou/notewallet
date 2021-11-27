package order_test

import (
	"encoding/json"
	"log"
	"myapp/Model/api"
	"myapp/Model/transaction"
	"myapp/Model/user"
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

var (
	userJSON = `{"userId":2,"coin":"bitcoin","jumlah_beli":10000}`
)

func TestBuyCritpo(t *testing.T) {

	e := InitEcho()
	baserespon := api.BaseResponse{}
	var user user.User
	var buy transaction.Transaction
	t.Run("Success Create user", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		err := configs.DB.First(&user, buy.UserId).Error
		log.Println(rec)
		c := e.NewContext(req, rec)
		c.SetPath("/crypto/buy")
		if user.ID == 0 {
			assert.Equal(t, http.StatusBadRequest, baserespon.Code)
		} else if err != nil {
			assert.Equal(t, http.StatusInternalServerError, baserespon.Code)
		}
		if assert.NoError(t, order.Pembelian(c)) {
			body := rec.Body.String()
			baseResponse := api.BaseResponse{}
			if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
				assert.Error(t, err, "Failed convert body to object")
			}
			assert.Equal(t, http.StatusOK, baseResponse.Code)
		}
	})
}
