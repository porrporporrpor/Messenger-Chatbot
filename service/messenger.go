package service

import (
	"errors"
	"fmt"
	"github.com/openlyinc/pointy"
	"github.com/rs/zerolog/log"
	"gitlab.com/pplayground/messenger-chatbot/backend/model"
	"gitlab.com/pplayground/messenger-chatbot/backend/util"
	"io/ioutil"
	"net/http"
)

type MessengerServiceInterface interface {
	CreateGetStartButton() error
	CreatePersistentMenu(psid string) error
	CreateQuickReply(psid string, quickReply model.QuickReplyMessage) error
	CreateGenericTemplate(psid string, payload model.TemplateGenericPayload) error
	CreateReceiptTemplate(psid string, payload model.TemplateReceiptPayload) error
	CreateButtonTemplate(psid string, payload model.TemplateButtonPayload) error
}

type MessengerService struct {
	Config model.Config
}

func (s MessengerService) CreateGetStartButton() error {
	endpoint := fmt.Sprintf("%s/messenger_profile?access_token=%s",
		s.Config.MessengerConfig.MessengerAPIUrl,
		s.Config.MessengerConfig.PageAccessToken)

	// Create Get Start Button
	rawRequestBody, err := util.CreateRequestBody(model.RequestBodyCreateGetStart{
		GetStart: model.GetStartPayload{Payload: "GET_START"},
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
				Text:   "???????????????????????????????????? {{user_first_name}}!",
			},
			{
				Locate: "en_US",
				Text:   "Hello {{user_first_name}}!",
			},
			{
				Locate: "th_TH",
				Text:   "?????????????????? {{user_first_name}}!",
			},
			{
				Locate: "zh_CN",
				Text:   "?????? {{user_first_name}}!",
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

func (s MessengerService) CreatePersistentMenu(psid string) error {
	endpoint := fmt.Sprintf("%s/custom_user_settings?access_token=%s",
		s.Config.MessengerConfig.MessengerAPIUrl,
		s.Config.MessengerConfig.PageAccessToken)

	rawRequestBody, err := util.CreateRequestBody(model.RequestBodyCreatePersistentMenu{
		PSID: psid,
		PersistentMenus: []model.PersistentMenu{
			{
				Locale:                "default",
				ComposerInputDisabled: false,
				CallToActions: []model.CallToAction{
					{
						Type:    "postback",
						Title:   pointy.String("Show Now !"),
						Payload: pointy.String("SHOP_NOW"),
					},
					{
						Type:               "web_url",
						Title:              pointy.String("Instagram"),
						URL:                pointy.String("https://www.instagram.com/threemandown"),
						WebviewHeightRatio: pointy.String("full"),
					},
					{
						Type:    "postback",
						Title:   pointy.String("My Order !"),
						Payload: pointy.String("MY_ORDER"),
					},
				},
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

	resp, err := http.Post(endpoint, model.ContentTypeJSON, rawRequestBody)
	if err != nil {
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

		log.Error().
			Str("type", model.LogTypeService).
			Str("status", model.LogStatusFailed).
			Msg(err.Error())
		return err
	}
	if resp.StatusCode != http.StatusOK {
		errorMessage := errors.New(fmt.Sprintf("call http POST given %s resposne", resp.Status))
		log.Error().
			Str("type", model.LogTypeService).
			Str("status", model.LogStatusFailed).
			Msg(errorMessage.Error())
		return errorMessage
	}

	return nil
}

func (s MessengerService) CreateQuickReply(psid string, quickReply model.QuickReplyMessage) error {
	endpoint := fmt.Sprintf("%s/messages?access_token=%s",
		s.Config.MessengerConfig.MessengerAPIUrl,
		s.Config.MessengerConfig.PageAccessToken)

	rawRequestBody, err := util.CreateRequestBody(model.RequestBodyCreateQuickReply{
		Recipient:   model.Recipient{ID: psid},
		MessageType: "RESPONSE",
		Message: model.QuickReplyMessage{
			Text:         quickReply.Text,
			QuickReplies: quickReply.QuickReplies,
		}})
	if err != nil {
		log.Error().
			Str("type", model.LogTypeService).
			Str("status", model.LogStatusFailed).
			Msg(err.Error())
		return err
	}

	resp, err := http.Post(endpoint, model.ContentTypeJSON, rawRequestBody)
	if err != nil {
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

		log.Error().
			Str("type", model.LogTypeService).
			Str("status", model.LogStatusFailed).
			Msg(err.Error())
		return err
	}
	if resp.StatusCode != http.StatusOK {
		errorMessage := errors.New(fmt.Sprintf("call http POST given %s resposne", resp.Status))
		log.Error().
			Str("type", model.LogTypeService).
			Str("status", model.LogStatusFailed).
			Msg(errorMessage.Error())
		return errorMessage
	}

	return nil
}

func (s MessengerService) CreateGenericTemplate(psid string, payload model.TemplateGenericPayload) error {
	endpoint := fmt.Sprintf("%s/messages?access_token=%s",
		s.Config.MessengerConfig.MessengerAPIUrl,
		s.Config.MessengerConfig.PageAccessToken)

	rawRequestBody, err := util.CreateRequestBody(model.RequestBodyCreateGenericTemplate{
		Recipient: model.Recipient{ID: psid},
		Message: model.TemplateGenericMessage{
			Attachment: model.TemplateGenericAttachment{
				Type:    "template",
				Payload: payload,
			}},
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

	return nil
}

func (s MessengerService) CreateReceiptTemplate(psid string, payload model.TemplateReceiptPayload) error {
	endpoint := fmt.Sprintf("%s/messages?access_token=%s",
		s.Config.MessengerConfig.MessengerAPIUrl,
		s.Config.MessengerConfig.PageAccessToken)

	rawRequestBody, err := util.CreateRequestBody(model.RequestBodyCreateReceiptTemplate{
		Recipient: model.Recipient{ID: psid},
		Message: model.TemplateReceiptMessage{
			Attachment: model.TemplateReceiptAttachment{
				Type:    "template",
				Payload: payload,
			}},
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

	return nil
}

func (s MessengerService) CreateButtonTemplate(psid string, payload model.TemplateButtonPayload) error {
	endpoint := fmt.Sprintf("%s/messages?access_token=%s",
		s.Config.MessengerConfig.MessengerAPIUrl,
		s.Config.MessengerConfig.PageAccessToken)

	rawRequestBody, err := util.CreateRequestBody(model.RequestBodyCreateButtonTemplate{
		Recipient: model.Recipient{ID: psid},
		Message: model.TemplateButtonMessage{Attachment: model.TemplateButtonAttachment{
			Type:    "template",
			Payload: payload,
		}},
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

	return nil
}
