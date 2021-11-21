package client

import (
	"github.com/gin-gonic/gin"
	"log"
	"mri/whatsapp-client-message/handler/service"
	"mri/whatsapp-client-message/models"
)

func (cl *ClientHandlerWhatsapp) SendMessage(c *gin.Context) {
	var param models.ParamSendMessage
	errRequest := c.Bind(&param)
	if errRequest != nil {
		c.JSON(400, models.CommonResponse{
			Code:      400,
			IsSuccess: false,
			Message:   "Parameter can't empty " + errRequest.Error(),
		})
		return
	}

	apiClientWa := service.NewWhatsappClientHandler()
	go func() {
		_, err := apiClientWa.SendMessage(param.Msisdn, param.Message)
		if err != nil {
			log.Println("Info Error:", err.Error())
		}
	}()

	c.JSON(200, models.CommonResponse{
		Code:      200,
		IsSuccess: true,
		Message:   "Success send message to " + param.Msisdn,
		Data:      nil,
	})
}

func (cl *ClientHandlerWhatsapp) SendDocumentMessage(c *gin.Context) {
	var param models.ParamSendDocument
	errRequest := c.Bind(&param)
	if errRequest != nil {
		c.JSON(400, models.CommonResponse{
			Code:      400,
			IsSuccess: false,
			Message:   "Parameter can't empty " + errRequest.Error(),
		})
		return
	}

	apiClientWa := service.NewWhatsappClientHandler()
	go func() {
		_, err := apiClientWa.SendMessageWithDocument(param.Msisdn, param.Message, param.DocumentLink)
		if err != nil {
			log.Println("Info Error:", err.Error())
		}
	}()

	c.JSON(200, models.CommonResponse{
		Code:      200,
		IsSuccess: true,
		Message:   "Success send message to " + param.Msisdn,
		Data:      nil,
	})
}
