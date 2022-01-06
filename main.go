package main

import (
	"fmt"
	"log"
	"mri/whatsapp-client-message/db"
	"mri/whatsapp-client-message/handler"
	"mri/whatsapp-client-message/handler/client"
	"mri/whatsapp-client-message/handler/service"
	"mri/whatsapp-client-message/models"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db.InitConnectionFromEnvirontment().CreateNewConnection()
	db.Connection.AutoMigrate(&models.LogWhatsapp{})
	db.Connection.AutoMigrate(&models.Config{})
	server := handler.RunServer()

	handlerClient := client.NewClientHandlerWhatsapp()
	api := server.Group("api/v1/whatsapp")
	{
		api.POST("/send-notification-message", handlerClient.SendMessage)
		api.POST("/send-notification-document", handlerClient.SendDocumentMessage)

		api.GET("/logs", handlerClient.GetLogWhatsapp)
	}

	// resendFailed()
	server.Run()
}

func resendFailed()  {
	pingTicker := time.NewTicker(30 * time.Minute)
	pingDone := make(chan bool)
	go func() {
		for {
			select {
			case <-pingDone:
				return
			case <-pingTicker.C:
				fmt.Println("INSTANCE RESEND")
				resend := resend()
				if !resend {
					pingDone <- true
				}
			}
		}
	}()
}

func resend() bool {
	var config models.Config
	db.Connection.Where("application_code = ? AND need_resend = ? AND services = ?", "budget-001", true, "wa").First(&config)

	if config.NeedResend {
		logFaileds := getWhatsappFaileds()
		for _, logFailed := range logFaileds {
			fmt.Println("RESEND", logFailed.Id, logFailed.To)
			apiClientWa := service.NewWhatsappClientNusaGateWayHandler()
			_, err := apiClientWa.ReSendMessage(logFailed.To, logFailed.Message, logFailed.Id)
			if err != nil {
				log.Println("Info Error:", err.Error())
			}
		}
	}

	return true
}

func getWhatsappFaileds() (faileds []models.LogWhatsapp) {
	db.Connection.Where("status = ? AND count_resend <= ?", false, 3, nil).Limit(5).Order("id asc").Find(&faileds)
	return
}
