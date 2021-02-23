package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/pplayground/messenger-chatbot/backend/model"
	"os"
)

func Logger(config model.Config) gin.HandlerFunc {
	if config.ENVState != "develop" {
		return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			log.Logger = zerolog.New(os.Stderr).With().Caller().Timestamp().Logger()
			log.Debug().
				Str("type", model.LogTypeMiddleware).
				Str("status", model.LogStatusRequest).
				Msg(fmt.Sprintf("%d | %s | %s | %s %s",
					param.StatusCode,
					param.Latency,
					param.ClientIP,
					param.Method,
					param.Path))
			return ""
		})
	} else {
		return gin.Logger()
	}
}
