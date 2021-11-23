package createuser

import (
	"fmt"

	"myapp/Model/api"
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
	// if err:=configs.DB.Where("name=? AND email=?",user.Name,user.Email).First(&user).Error; err!=nil{
	// 	return c.JSON(http.StatusBadRequest, BaseResponse{
	// 		http.StatusInternalServerError,
	// 		"salah",
	// 		nil,
	// 	})
	// }
	_, err := mail.ParseAddress(user.Email)
	if err == nil && user.Name != "" {
		fmt.Println("acces open")
		result := configs.DB.Create(&user)
		Token = middlewares.GenerateTokenJWT(int(user.ID))
		//jwt
		if result.Error != nil {

			return c.JSON(http.StatusBadRequest, api.BaseResponse{
				http.StatusInternalServerError,
				result.Error.Error(),
				nil,
			})
		}
	} else {
		falseInput := []string{}
		if user.Name == "" {
			falseInput = append(falseInput, "username kosong")
		}
		falseInput = append(falseInput, "wrong email")
		return c.JSON(http.StatusNotAcceptable, falseInput)
	}
	var response = api.BaseResponse{
		http.StatusOK,
		"sukses",
		map[string]interface{}{
			"token": Token,
		},
	}
	return c.JSON(http.StatusOK, response)
}
