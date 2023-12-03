package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/playwright-community/playwright-go"
	"github.com/seew0/loginAutomation/api"
)

func init() {
	err := playwright.Install()
	if err != nil {
		log.Fatalf("Error installing Playwright: %v", err)
	}
}

func main() {
	router := gin.Default()
	Server := api.NewServer(":4000", router)
	Server.Start()
}
