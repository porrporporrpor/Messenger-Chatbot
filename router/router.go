package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/pplayground/messenger-chatbot/backend/api"
	"gitlab.com/pplayground/messenger-chatbot/backend/middleware"
	"gitlab.com/pplayground/messenger-chatbot/backend/model"
)

func webhookAPIGroup(r *gin.RouterGroup) {
	r.GET("", api.VerifyWebhook)
	r.POST("", api.ReceiveFromWebhook)
}

func Start(r *gin.Engine, config model.Config) {
	r.Use(middleware.Logger(config))

	webhookAPIGroup(r.Group("/webhook"))

	r.Run(fmt.Sprintf(":%v", config.Port))
}
