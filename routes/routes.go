package routes

import (
	createuser "myapp/controler/createUser"
	"myapp/controler/crypto"
	"myapp/controler/order"
	middlewares "myapp/midleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	config := middleware.JWTConfig{
		Claims:     &middlewares.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	e := echo.New()
	middlewares.Middlewares(e)
	e.Pre(middleware.RemoveTrailingSlash())
	eUser := e.Group("/users")
	eCrypto := e.Group("/crypto")
	eTransactions := e.Group("/transactions")
	eCrypto.Use(middleware.JWTWithConfig(config))
	eTransactions.Use(middleware.JWTWithConfig(config))
	eUser.POST("/signup", createuser.CreateUser)
	eCrypto.GET("/price/:id", crypto.Cekprice)
	eTransactions.POST("/buy", order.Pembelian)
	eTransactions.POST("/sell", order.Penjualan)
	eTransactions.GET("/result/:id", order.Result)
	return e
}
