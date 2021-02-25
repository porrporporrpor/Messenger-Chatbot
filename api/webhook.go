package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gitlab.com/pplayground/messenger-chatbot/backend/model"
	"gitlab.com/pplayground/messenger-chatbot/backend/service"
	"net/http"
	"strconv"
)

type WebhookAPI struct {
	MessengerService service.MessengerServiceInterface
}

func (api WebhookAPI) VerifyWebhook(ctx *gin.Context) {
	challenge := ctx.Query("hub.challenge")
	intChallenge, _ := strconv.Atoi(challenge)
	ctx.JSON(200, intChallenge)
}

func (api WebhookAPI) ReceiveFromWebhook(ctx *gin.Context) {
	var req model.MessengerRequestBody
	ctx.ShouldBindJSON(&req)

	messaging := (*req.Entry[0].Messaging)[0]
	if messaging.Message != nil {
		if messaging.Message.Text != "" {
			switch messaging.Message.Text {
			case model.TextOperationGetStart:
				err := api.MessengerService.CreatePersistentMenu(messaging.Sender.ID)
				if err != nil {
					log.Error().
						Str("type", model.LogTypeAPI).
						Str("status", model.LogStatusFailed).
						Msg("operation CreatePersistentMenu service failed")
					ctx.JSON(http.StatusInternalServerError, model.ResponsePayload{
						Status:  model.LogStatusFailed,
						Payload: "cannot create persistent menu"})
				}
			}
		}
	}
	if messaging.Postback != nil {
		fmt.Println(messaging.Postback)
		switch messaging.Postback.Payload {
		case "SHOP_NOW":
			api.MessengerService.CreateShopNowTemplate(messaging.Sender.ID)
		case "MY_ORDER":
			fmt.Println("my order is not implement")
		}

	}
	ctx.JSON(200, nil)
}
