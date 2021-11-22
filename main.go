package main

import (
	"github.com/joho/godotenv"
	"mri/whatsapp-client-message/db"
	"mri/whatsapp-client-message/handler"
	"mri/whatsapp-client-message/handler/client"
	"mri/whatsapp-client-message/models"
)

func main() {
	_ = godotenv.Load()
	db.InitConnectionFromEnvirontment().CreateNewConnection()
	db.Connection.AutoMigrate(&models.LogWhatsapp{})
	server := handler.RunServer()

	handlerClient := client.NewClientHandlerWhatsapp()
	api := server.Group("api/v1/whatsapp")
	{
		api.POST("/send-notification-message", handlerClient.SendMessage)
		api.POST("/send-notification-document", handlerClient.SendDocumentMessage)
	}

	server.Run()
}
