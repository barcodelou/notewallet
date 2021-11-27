package createuser

import (
	"encoding/json"
	"log"
	"myapp/Model/api"

	configs "myapp/config"
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
	userJSON   = `{"nama":"lor","email":"jona@labstack.com","asset":0}`
	BlankName  = `{"nama":"","email":"jona@labstack.com","asset":0}`
	FalseEmail = `{"nama":"lor","email":"jonalabstack.com","asset":0}`
)

func TestCreateUserControllers(t *testing.T) {
	e := InitEcho()
	// CreateSeedUser()
	t.Run("Success Create user", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		log.Println(rec)
		c := e.NewContext(req, rec)
		c.SetPath("/users/signup")

		if assert.NoError(t, CreateUser(c)) {
			body := rec.Body.String()
			baseResponse := api.BaseResponse{}
			if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
				assert.Error(t, err, "Failed convert body to object")
			}
			assert.Equal(t, http.StatusOK, baseResponse.Code)
		}
	})
	t.Run("blank name false", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(BlankName))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		log.Println(rec)
		c := e.NewContext(req, rec)
		c.SetPath("/users/signup")

		if assert.NoError(t, CreateUser(c)) {
			body := rec.Body.String()
			baseResponse := api.BaseResponse{}
			if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
				assert.Error(t, err, "Failed convert body to object")
			}
			assert.Equal(t, http.StatusNotAcceptable, baseResponse.Code)
		}
	})

}
