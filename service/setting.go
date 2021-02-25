package service

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"gitlab.com/pplayground/messenger-chatbot/backend/model"
	"gitlab.com/pplayground/messenger-chatbot/backend/util"
	"io/ioutil"
	"net/http"
)

type SettingServiceInterface interface {
	CreateGetStartButton() error
}

type SettingService struct {
	Config model.Config
}

func (s SettingService) CreateGetStartButton() error {
	endpoint := fmt.Sprintf("%s/messenger_profile?access_token=%s",
		s.Config.MessengerConfig.MessengerAPIUrl,
		s.Config.MessengerConfig.PageAccessToken)

	// Create Get Start Button
	rawRequestBody, err := util.CreateRequestBody(model.RequestBodyCreateGetStart{
		GetStart: model.GetStartPayload{Payload: "<postback_payload>"},
	})
	if err != nil {
		log.Error().
			Str("type", model.LogTypeService).
			Str("status", model.LogStatusFailed).
			Msg(err.Error())
		return err
	}

	resp, err := http.Post(endpoint, model.ContentTypeJSON, rawRequestBody)
	if err != nil {
		log.Error().
			Str("type", model.LogTypeService).
			Str("status", model.LogStatusFailed).
			Msg(err.Error())
		return err
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error().
				Str("type", model.LogTypeService).
				Str("status", model.LogStatusFailed).
				Msg(err.Error())
			return err
		}
		bodyString := string(bodyBytes)
		log.Debug().
			Str("type", model.LogTypeService).
			Str("status", model.LogStatusData).
			Msg(bodyString)

		errorMessage := errors.New(fmt.Sprintf("call http POST given %s resposne", resp.Status))
		log.Error().
			Str("type", model.LogTypeService).
			Str("status", model.LogStatusFailed).
			Msg(errorMessage.Error())
		return errorMessage
	}

	// Create Greeting Message
	rawRequestBody, err = util.CreateRequestBody(model.RequestBodyCreateGreetingMessage{
		GreetingMessages: []model.GreetingMessage{
			{
				Locate: "default",
				Text:   "สวัสดีจ้าาาา {{user_first_name}}!",
			},
			{
				Locate: "en_US",
				Text:   "Hello {{user_first_name}}!",
			},
			{
				Locate: "th_TH",
				Text:   "สวัสดี {{user_first_name}}!",
			},
			{
				Locate: "zh_CN",
				Text:   "你好 {{user_first_name}}!",
			},
		},
	})
	if err != nil {
		log.Error().
			Str("type", model.LogTypeService).
			Str("status", model.LogStatusFailed).
			Msg(err.Error())
		return err
	}

	resp, err = http.Post(endpoint, model.ContentTypeJSON, rawRequestBody)
	if err != nil {
		log.Error().
			Str("type", model.LogTypeService).
			Str("status", model.LogStatusFailed).
			Msg(err.Error())
		return err
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error().
				Str("type", model.LogTypeService).
				Str("status", model.LogStatusFailed).
				Msg(err.Error())
			return err
		}
		bodyString := string(bodyBytes)
		log.Debug().
			Str("type", model.LogTypeService).
			Str("status", model.LogStatusData).
			Msg(bodyString)

		errorMessage := errors.New(fmt.Sprintf("call http POST given %s resposne", resp.Status))
		log.Error().
			Str("type", model.LogTypeService).
			Str("status", model.LogStatusFailed).
			Msg(errorMessage.Error())
		return errorMessage
	}

	return nil
}
