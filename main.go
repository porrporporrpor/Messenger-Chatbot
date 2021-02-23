package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/pplayground/messenger-chatbot/backend/model"
	"gitlab.com/pplayground/messenger-chatbot/backend/router"
)

func main() {
	_ = godotenv.Load()

	config := model.Config{}
	err := config.Load()
	if err != nil {
		panic(err)
	}

	if config.ENVMode != "develop" {
		gin.SetMode(gin.ReleaseMode)
	}
	router.Start(gin.Default(), config)
}
