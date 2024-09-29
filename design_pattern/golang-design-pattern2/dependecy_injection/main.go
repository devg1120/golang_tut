package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	config := newConfig()
	h, err := initializeHandlers(config)
	if err != nil {
		log.Fatal(err)
	}
	e.PUT("/withdraw", h.BankHandler.Withdraw)
	e.Logger.Fatal(e.Start(":8080"))
}
