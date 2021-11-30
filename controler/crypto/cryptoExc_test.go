package crypto

import (
	configs "myapp/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	configs.InitConfigDBTest()
	e := echo.New()
	return e
}

// type (
// 	Kripto struct {
// 		Code    int     `json:"code"`
// 		Message string  `json:"message"`
// 		Harga   float32 `json:"harga"`
// 		Uang    string  `json:"uang"`
// 	}

// 	handler struct {
// 		db map[string]*Kripto
// 	}
// )

// var (
// 	mockDB = map[string]*Kripto{
// 		"bitcoin": &Kripto{200, "bitcoin", 100000, "idr"},
// 	}
// 	userJSON = `{"code":200,"message":"bitcoin","harga":100000,uang":"idr"}`
// )

func TestGetUser(t *testing.T) {

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/crypto/price/:id")
	c.SetParamNames("id")
	c.SetParamValues("bitcoin")

	// Assertions
	if assert.NoError(t, Cekprice(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
