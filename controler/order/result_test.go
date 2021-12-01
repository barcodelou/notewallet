package order_test

import (
	"myapp/Model/result"
	"myapp/controler/order"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func CreateSeedNews() {
// 	var result result.Result
// 	result.UserId = 2
// 	result.AssetResult = 15000
// 	result.UserName = "yonathan"
// 	result.Conclusion = "profit"
// 	configs.DB.Create(&result)
// }

func TestResult(t *testing.T) {
	// Setup
	var result result.Result
	e := InitEcho()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/transaction/result/:id")
	c.SetParamNames("id")
	c.SetParamValues("3")

	// Assertions
	if assert.NoError(t, order.Result(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotNil(t, result)
	}
	assert.NotEqual(t, http.StatusBadRequest, rec.Code)
}
