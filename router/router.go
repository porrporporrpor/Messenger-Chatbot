package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/pplayground/messenger-chatbot/backend/api"
	"gitlab.com/pplayground/messenger-chatbot/backend/model"
)

func webhookAPIGroup(r *gin.RouterGroup) {
	r.GET("", api.VerifyWebhook)
	r.POST("", api.ReceiveFromWebhook)
}

func Start(r *gin.Engine, config model.Config) {
	webhookAPIGroup(r.Group("/webhook"))

	r.Run(fmt.Sprintf(":%v", config.Port))
}
