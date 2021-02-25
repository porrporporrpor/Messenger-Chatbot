package model

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type MessengerConfig struct {
	MessengerAPIUrl string
	AppSecret       string
	PageAccessToken string
	ValidationToken string
}

type Config struct {
	MessengerConfig MessengerConfig
	ENVState        string
	Port            string
}

func (config *Config) Load() error {
	var missingEnv []string

	envState := os.Getenv("ENV_STATE")
	if envState == "" {
		config.ENVState = "develop"
	} else {
		config.ENVState = envState
	}

	port := os.Getenv("PORT")
	if port == "" {
		config.Port = "8080"
	} else {
		config.Port = port
	}

	appSecret := os.Getenv("APP_SECRET")
	if appSecret == "" {
		missingEnv = append(missingEnv, "APP_SECRET")
	}

	messengerAPIUrl := os.Getenv("MESSENGER_API_URL")
	if messengerAPIUrl == "" {
		missingEnv = append(missingEnv, "MESSENGER_API_URL")
	}

	pageAccessToken := os.Getenv("PAGE_ACCESS_TOKEN")
	if pageAccessToken == "" {
		missingEnv = append(missingEnv, "PAGE_ACCESS_TOKEN")
	}

	validationToken := os.Getenv("VALIDATION_TOKEN")
	if validationToken == "" {
		missingEnv = append(missingEnv, "VALIDATION_TOKEN")
	}

	if len(missingEnv) > 0 {
		return errors.New(fmt.Sprintf("Environment is required %v", strings.Join(missingEnv, ", ")))
	}
	config.MessengerConfig.MessengerAPIUrl = messengerAPIUrl
	config.MessengerConfig.AppSecret = appSecret
	config.MessengerConfig.PageAccessToken = pageAccessToken
	config.MessengerConfig.ValidationToken = validationToken
	return nil
}
