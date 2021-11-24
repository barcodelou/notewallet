package createuser

import (
	"fmt"

	"myapp/Model/api"
	"myapp/Model/handlerfalse"
	"myapp/Model/user"
	configs "myapp/config"
	middlewares "myapp/midleware"
	"net/http"
	"net/mail"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	var Token string
	var user user.User
	c.Bind(&user)
	_, err := mail.ParseAddress(user.Email)
	if err == nil && user.Name != "" {
		fmt.Println("acces open")
		result := configs.DB.Create(&user)
		Token = middlewares.GenerateTokenJWT(int(user.ID))
		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, api.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: result.Error.Error(),
				Data:    nil,
			})
		}
	} else {
		return c.JSON(http.StatusNotAcceptable, handlerfalse.HandlerUser(user.Name, err))
	}
	var response = api.BaseResponse{
		Code:    http.StatusOK,
		Message: "sukses",
		Data: map[string]interface{}{
			"token": Token,
		},
	}
	return c.JSON(http.StatusOK, response)
}
