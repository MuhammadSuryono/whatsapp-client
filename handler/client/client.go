package client

import "github.com/gin-gonic/gin"

type IClientHandlerWhatsapp interface {
	SendMessage(c *gin.Context)
}

type ClientHandlerWhatsapp struct {

}

func NewClientHandlerWhatsapp() IClientHandlerWhatsapp {
	return &ClientHandlerWhatsapp{}
}
