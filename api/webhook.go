package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gitlab.com/pplayground/messenger-chatbot/backend/model"
	"gitlab.com/pplayground/messenger-chatbot/backend/service"
	"net/http"
	"strconv"
	"strings"
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

			}
		}
	}
	if messaging.Postback != nil {
		log.Debug().
			Str("type", model.LogTypeAPI).
			Str("status", model.LogStatusData).
			Msg(fmt.Sprintf("postback : %v", messaging.Postback))
		switch messaging.Postback.Payload {
		case "GET_START":
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
		case "SHOP_NOW":
			err := api.MessengerService.CreateShopNowTemplate(messaging.Sender.ID)
			if err != nil {
				log.Error().
					Str("type", model.LogTypeAPI).
					Str("status", model.LogStatusFailed).
					Msg("operation CreateShopNowTemplate service failed")
				ctx.JSON(http.StatusInternalServerError, model.ResponsePayload{
					Status:  model.LogStatusFailed,
					Payload: "cannot create shop now template menu"})
			}
		case "MY_ORDER":
			log.Debug().
				Str("type", model.LogTypeAPI).
				Str("status", model.LogStatusData).
				Msg("my order is not implement")
		default:
			if strings.Contains(messaging.Postback.Payload, "VIEW_PRODUCT") {
				log.Debug().
					Str("type", model.LogTypeAPI).
					Str("status", model.LogStatusData).
					Msg(fmt.Sprintf("postback payload : %v", messaging.Postback.Payload))
			}
		}
	}
	ctx.JSON(200, nil)
}
