package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/openlyinc/pointy"
	"github.com/rs/zerolog/log"
	"gitlab.com/pplayground/messenger-chatbot/backend/model"
	"gitlab.com/pplayground/messenger-chatbot/backend/service"
	"net/http"
	"strconv"
	"strings"
)

type WebhookAPI struct {
	MessengerService service.MessengerServiceInterface
}

func (api WebhookAPI) VerifyWebhook(ctx *gin.Context) {
	challenge := ctx.Query("hub.challenge")
	intChallenge, _ := strconv.Atoi(challenge)
	ctx.JSON(200, intChallenge)
}

func (api WebhookAPI) ReceiveFromWebhook(ctx *gin.Context) {
	var req model.MessengerRequestBody
	ctx.ShouldBindJSON(&req)

	messaging := (*req.Entry[0].Messaging)[0]
	if messaging.Message != nil {
		log.Debug().
			Str("type", model.LogTypeAPI).
			Str("status", model.LogStatusData).
			Msg(fmt.Sprintf("message : %v", messaging.Message))
		if messaging.Message.QuickReply != nil {
			switch messaging.Message.QuickReply.Payload {
			case "SHOP_NOW":
				err := api.MessengerService.CreateGenericTemplate(messaging.Sender.ID, model.TemplateGenericPayload{
					TemplateType: "generic",
					Elements: []model.GenericElement{
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
				})
				if err != nil {
					log.Error().
						Str("type", model.LogTypeAPI).
						Str("status", model.LogStatusFailed).
						Msg("operation CreateShopNowTemplate service failed")
					ctx.JSON(http.StatusInternalServerError, model.ResponsePayload{
						Status:  model.LogStatusFailed,
						Payload: "cannot create shop now template menu"})
				}
			case "MY_ORDER":
				err := api.MessengerService.CreateReceiptTemplate(messaging.Sender.ID, model.TemplateReceiptPayload{
					TemplateType:  "receipt",
					RecipientName: "Benjawan S.",
					OrderName:     "123456",
					Currency:      "THB",
					PaymentMethod: "Cash",
					OrderUrl:      "http://petersapparel.parseapp.com/order?order_id=123456",
					Timestamp:     "1428444852",
					Summary: model.ReceiptSummary{
						Subtotal:     750.00,
						ShippingCost: 40.95,
						TotalTax:     60.19,
						TotalCost:    460.14,
					},
					Adjustments: &[]model.ReceiptAdjustments{
						{
							Name:   "Discount",
							Amount: 100,
						},
					},
					Elements: []model.ReceiptElement{
						{
							Title:    "Classic White T-Shirt",
							Subtitle: "100% Soft and Luxurious Cotton",
							Quantity: 2,
							Price:    500,
							Currency: "THB",
							ImageUrl: "https://picsum.photos/200",
						},
						{
							Title:    "Classic Pink T-Shirt",
							Subtitle: "100% Soft and Luxurious Cotton",
							Quantity: 1,
							Price:    250,
							Currency: "THB",
							ImageUrl: "https://picsum.photos/200",
						},
					},
				})
				if err != nil {
					log.Error().
						Str("type", model.LogTypeAPI).
						Str("status", model.LogStatusFailed).
						Msg("operation CreateReceiptTemplate service failed")
					ctx.JSON(http.StatusInternalServerError, model.ResponsePayload{
						Status:  model.LogStatusFailed,
						Payload: "cannot create receipt template of my order menu"})
				}
			}
		}
	}
	if messaging.Postback != nil {
		log.Debug().
			Str("type", model.LogTypeAPI).
			Str("status", model.LogStatusData).
			Msg(fmt.Sprintf("postback : %v", messaging.Postback))
		switch messaging.Postback.Payload {
		case "GET_START":
			err := api.MessengerService.CreatePersistentMenu(messaging.Sender.ID)
			if err != nil {
				log.Error().
					Str("type", model.LogTypeAPI).
					Str("status", model.LogStatusFailed).
					Msg("operation CreatePersistentMenu service failed")
				ctx.JSON(http.StatusInternalServerError, model.ResponsePayload{
					Status:  model.LogStatusFailed,
					Payload: "cannot create persistent menu"})
			}

			quickReply := model.QuickReplyMessage{
				Text: "ขอบคุณที่สนใจรับบริการจากทางเราค่ะ โปรดบอกเราในสิ่งที่คุณต้องการได้เลย!",
				QuickReplies: []model.QuickReply{
					{
						ContentType: "text",
						Title:       "Show Now",
						Payload:     "SHOP_NOW",
					},
					{
						ContentType: "text",
						Title:       "My Order",
						Payload:     "MY_ORDER",
					},
				},
			}
			err = api.MessengerService.CreateQuickReply(messaging.Sender.ID, quickReply)
			if err != nil {
				log.Error().
					Str("type", model.LogTypeAPI).
					Str("status", model.LogStatusFailed).
					Msg("operation CreatePersistentMenu service failed")
				ctx.JSON(http.StatusInternalServerError, model.ResponsePayload{
					Status:  model.LogStatusFailed,
					Payload: "cannot create persistent menu"})
			}
		case "SHOP_NOW":
			err := api.MessengerService.CreateGenericTemplate(messaging.Sender.ID, model.TemplateGenericPayload{
				TemplateType: "generic",
				Elements: []model.GenericElement{
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
			})
			if err != nil {
				log.Error().
					Str("type", model.LogTypeAPI).
					Str("status", model.LogStatusFailed).
					Msg("operation CreateShopNowTemplate service failed")
				ctx.JSON(http.StatusInternalServerError, model.ResponsePayload{
					Status:  model.LogStatusFailed,
					Payload: "cannot create shop now template menu"})
			}
		case "MY_ORDER":
			err := api.MessengerService.CreateReceiptTemplate(messaging.Sender.ID, model.TemplateReceiptPayload{
				TemplateType:  "receipt",
				RecipientName: "Benjawan S.",
				OrderName:     "123456",
				Currency:      "THB",
				PaymentMethod: "Cash",
				OrderUrl:      "http://petersapparel.parseapp.com/order?order_id=123456",
				Timestamp:     "1428444852",
				Summary: model.ReceiptSummary{
					Subtotal:     750.00,
					ShippingCost: 40.95,
					TotalTax:     60.19,
					TotalCost:    460.14,
				},
				Adjustments: &[]model.ReceiptAdjustments{
					{
						Name:   "Discount",
						Amount: 100,
					},
				},
				Elements: []model.ReceiptElement{
					{
						Title:    "Classic White T-Shirt",
						Subtitle: "100% Soft and Luxurious Cotton",
						Quantity: 2,
						Price:    500,
						Currency: "THB",
						ImageUrl: "https://picsum.photos/200",
					},
					{
						Title:    "Classic Pink T-Shirt",
						Subtitle: "100% Soft and Luxurious Cotton",
						Quantity: 1,
						Price:    250,
						Currency: "THB",
						ImageUrl: "https://picsum.photos/200",
					},
				},
			})
			if err != nil {
				log.Error().
					Str("type", model.LogTypeAPI).
					Str("status", model.LogStatusFailed).
					Msg("operation CreateReceiptTemplate service failed")
				ctx.JSON(http.StatusInternalServerError, model.ResponsePayload{
					Status:  model.LogStatusFailed,
					Payload: "cannot create receipt template of my order menu"})
			}
		default:
			if strings.Contains(messaging.Postback.Payload, "VIEW_PRODUCT") {
				err := api.MessengerService.CreateButtonTemplate(messaging.Sender.ID, model.TemplateButtonPayload{
					TemplateType: "button",
					Text:         "กรุณาระบุจำนวนที่ต้องการสั่งซื้อ (เพิ่มเติม: สั่งมากกว่า 3 ชิ้นโปรดติดต่อแอดมิน)",
					Button: []model.CallToAction{
						{
							Type:    "postback",
							Title:   pointy.String("1"),
							Payload: pointy.String("QTY_1"),
						},
						{
							Type:    "postback",
							Title:   pointy.String("2"),
							Payload: pointy.String("QTY_2"),
						},
						{
							Type:    "postback",
							Title:   pointy.String("3"),
							Payload: pointy.String("QTY_3"),
						},
					},
				})
				if err != nil {
					log.Error().
						Str("type", model.LogTypeAPI).
						Str("status", model.LogStatusFailed).
						Msg("operation CreateButtonTemplate service failed")
					ctx.JSON(http.StatusInternalServerError, model.ResponsePayload{
						Status:  model.LogStatusFailed,
						Payload: "cannot create button template"})
				}
			}
			if strings.Contains(messaging.Postback.Payload, "QTY") {
				err := api.MessengerService.CreateReceiptTemplate(messaging.Sender.ID, model.TemplateReceiptPayload{
					TemplateType:  "receipt",
					RecipientName: "Benjawan S.",
					OrderName:     "123456",
					Currency:      "THB",
					PaymentMethod: "Cash",
					OrderUrl:      "http://petersapparel.parseapp.com/order?order_id=123456",
					Timestamp:     "1428444852",
					Summary: model.ReceiptSummary{
						Subtotal:     750.00,
						ShippingCost: 40.95,
						TotalTax:     60.19,
						TotalCost:    460.14,
					},
					Adjustments: &[]model.ReceiptAdjustments{
						{
							Name:   "Discount",
							Amount: 100,
						},
					},
					Elements: []model.ReceiptElement{
						{
							Title:    "Classic White T-Shirt",
							Subtitle: "100% Soft and Luxurious Cotton",
							Quantity: 2,
							Price:    500,
							Currency: "THB",
							ImageUrl: "https://picsum.photos/200",
						},
						{
							Title:    "Classic Pink T-Shirt",
							Subtitle: "100% Soft and Luxurious Cotton",
							Quantity: 1,
							Price:    250,
							Currency: "THB",
							ImageUrl: "https://picsum.photos/200",
						},
					},
				})
				if err != nil {
					log.Error().
						Str("type", model.LogTypeAPI).
						Str("status", model.LogStatusFailed).
						Msg("operation CreateReceiptTemplate service failed")
					ctx.JSON(http.StatusInternalServerError, model.ResponsePayload{
						Status:  model.LogStatusFailed,
						Payload: "cannot create receipt template of my order menu"})
				}
			}
		}
	}
	ctx.JSON(200, nil)
}
