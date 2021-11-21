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
	isSended, err := apiClientWa.SendMessage(param.Msisdn, param.Message)
	if err != nil {
		log.Println("Info Error:", err.Error())
	}

	if isSended {
		c.JSON(200, models.CommonResponse{
			Code:      200,
			IsSuccess: true,
			Message:   "Success send message to " + param.Msisdn,
			Data:      nil,
		})
		return
	}

	c.JSON(500, models.CommonResponse{
		Code:      500,
		IsSuccess: true,
		Message:   "Failed send message to " + param.Msisdn,
		Data:      nil,
	})
}
