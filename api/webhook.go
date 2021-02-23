package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/pplayground/messenger-chatbot/backend/model"
	"strconv"
)

func VerifyWebhook(ctx *gin.Context) {
	challenge := ctx.Query("hub.challenge")
	intChallenge, _ := strconv.Atoi(challenge)
	ctx.JSON(200, intChallenge)
}

func ReceiveFromWebhook(ctx *gin.Context) {
	var req model.MessengerRequestBody
	ctx.ShouldBindJSON(&req)
	fmt.Println(req)
	ctx.JSON(200, nil)
}
