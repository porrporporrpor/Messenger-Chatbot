package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gitlab.com/pplayground/messenger-chatbot/backend/model"
	"gitlab.com/pplayground/messenger-chatbot/backend/service"
	"net/http"
)

type SettingAPI struct {
	MessengerService service.MessengerServiceInterface
}

func (api SettingAPI) CreateGetStartSetting(ctx *gin.Context) {
	err := api.MessengerService.CreateGetStartButton()
	if err != nil {
		log.Error().
			Str("type", model.LogTypeAPI).
			Str("status", model.LogStatusFailed).
			Msg("operation CreateGetStartSetting service failed")
		ctx.JSON(http.StatusInternalServerError, model.ResponsePayload{
			Status:  model.LogStatusFailed,
			Payload: "cannot create get start setting"})
	}

	ctx.JSON(http.StatusOK, model.ResponsePayload{Status: model.LogStatusSuccess})
}
