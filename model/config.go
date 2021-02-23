package model

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type MessengerConfig struct {
	AppSecret       string
	PageAccessToken string
	ValidationToken string
}

type Config struct {
	MessengerConfig MessengerConfig
	Port            string
}

func (config *Config) Load() error {
	var missingEnv []string

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
	config.MessengerConfig.AppSecret = appSecret
	config.MessengerConfig.PageAccessToken = pageAccessToken
	config.MessengerConfig.ValidationToken = validationToken
	return nil
}
