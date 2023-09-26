package main

import (
	"FinalProject/configs"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	var config = configs.InitConfig()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)))
}
