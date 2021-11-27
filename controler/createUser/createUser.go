package createuser

import (
	"myapp/Model/api"
	"myapp/Model/user"
	configs "myapp/config"
	handlerespon "myapp/controler/order/handleRespon"
	middlewares "myapp/midleware"
	"net/http"
	"net/mail"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	var Token string
	var user user.User
	c.Bind(&user)
	err := Email(user.Email)
	if err == nil && user.Name != "" {
		result := configs.DB.Create(&user)
		Token = middlewares.GenerateTokenJWT(int(user.ID))
		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, api.BaseResponse{
				Code:    http.StatusBadRequest,
				Message: result.Error.Error(),
				Data:    nil,
			})
		}
	} else {
		return c.JSON(http.StatusNotAcceptable,
			api.BaseResponse{
				Code: http.StatusNotAcceptable,
				Data: handlerespon.HandlerUser(user.Name, err),
			})
	}
	return c.JSON(http.StatusOK, handlerespon.HandlerGetJwt(api.BaseResponse{}, Token))
}

func Email(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}
