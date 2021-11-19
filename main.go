package main

import (
	"github.com/joho/godotenv"
	"mri/whatsapp-client-message/handler"
	"mri/whatsapp-client-message/handler/client"
)

func main()  {
	_= godotenv.Load()
	server := handler.RunServer()

	handlerClient := client.NewClientHandlerWhatsapp()
	api := server.Group("api/v1/whatsapp")
	{
		api.POST("/send-notification", handlerClient.SendMessage)
	}

	server.Run()
}