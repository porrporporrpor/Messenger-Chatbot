package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/pplayground/messenger-chatbot/backend/api"
	"gitlab.com/pplayground/messenger-chatbot/backend/middleware"
	"gitlab.com/pplayground/messenger-chatbot/backend/model"
	"gitlab.com/pplayground/messenger-chatbot/backend/service"
	"os"
)

func webhookAPIGroup(r *gin.RouterGroup, config model.Config) {
	webhookAPI := api.WebhookAPI{MessengerService: service.MessengerService{Config: config}}

	r.GET("", webhookAPI.VerifyWebhook)
	r.POST("", webhookAPI.ReceiveFromWebhook)
}

func settingViewAPIGroup(r *gin.RouterGroup, config model.Config) {
	settingAPI := api.SettingAPI{SettingService: service.SettingService{Config: config}}

	r.POST("get_start", settingAPI.CreateGetStartSetting)
}

func Start(r *gin.Engine, config model.Config) {
	log.Logger = log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()
	if config.ENVState != "develop" {
		log.Logger = zerolog.New(os.Stderr).With().Caller().Timestamp().Logger()
	}

	r.Use(middleware.Logger(config))

	webhookAPIGroup(r.Group("/webhook"), config)
	settingViewAPIGroup(r.Group("/setting"), config)

	r.Run(fmt.Sprintf(":%v", config.Port))
}
