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
	CreatePersistentMenu(psid string) error
	CreateShopNowTemplate(psid string) error
}

type MessengerService struct {
	Config model.Config
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
						URL:                pointy.String("https://www.instagram.com/ppor.s"),
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

func (s MessengerService) CreateShopNowTemplate(psid string) error {
	endpoint := fmt.Sprintf("%s/messages?access_token=%s",
		s.Config.MessengerConfig.MessengerAPIUrl,
		s.Config.MessengerConfig.PageAccessToken)

	rawRequestBody, err := util.CreateRequestBody(model.RequestBodyCreateGenericTemplate{
		Recipient: model.Recipient{ID: psid},
		Message: model.TemplateMessage{
			Attachment: model.TemplateAttachment{
				Type: "template",
				Payload: model.TemplateAttachmentPayload{
					TemplateType: "generic",
					Elements: []model.Element{
						{
							Title:    "Product Name 1",
							ImageUrl: "https://picsum.photos/200",
							Subtitle: "Description",
							DefaultAction: model.CallToAction{
								Type:               "web_url",
								URL:                pointy.String("https://picsum.photos/200"),
								WebviewHeightRatio: pointy.String("tall"),
							},
							Buttons: []model.CallToAction{
								{
									Type:    "postback",
									Title:   pointy.String("Shop Now !"),
									Payload: pointy.String("VIEW_PRODUCT_1"),
								},
								{
									Type:  "web_url",
									Title: pointy.String("Instagram"),
									URL:   pointy.String("https://instagram.com"),
								},
							},
						},
						{
							Title:    "Product Name 2",
							ImageUrl: "https://picsum.photos/200",
							Subtitle: "Description",
							DefaultAction: model.CallToAction{
								Type:               "web_url",
								URL:                pointy.String("https://picsum.photos/200"),
								WebviewHeightRatio: pointy.String("tall"),
							},
							Buttons: []model.CallToAction{
								{
									Type:    "postback",
									Title:   pointy.String("Shop Now !"),
									Payload: pointy.String("VIEW_PRODUCT_2"),
								},
								{
									Type:  "web_url",
									Title: pointy.String("Instagram"),
									URL:   pointy.String("https://instagram.com"),
								},
							},
						},
					},
				},
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
