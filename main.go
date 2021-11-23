package main

import (
	configs "myapp/config"
	"myapp/routes"
)

func main() {
	configs.CnnectDB()
	e := routes.New()
	e.Start(":8000")
}
