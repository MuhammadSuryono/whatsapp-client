package service

import (
	"fmt"
	"log"
	"mri/whatsapp-client-message/db"
	"mri/whatsapp-client-message/models"
	"testing"

	"github.com/joho/godotenv"
)

func TestSendMessageWA(t *testing.T) {
	errLoadEnv := godotenv.Load()
	db.InitConnectionFromEnvirontment().CreateNewConnection()
	db.Connection.AutoMigrate(&models.LogWhatsapp{})
	if errLoadEnv != nil {
		log.Fatalln("Can't load env")
	}
	wa := NewWhatsappClientNusaGateWayHandler()
	status, err := wa.SendMessageWithDocument("0895355698652", "Test from api golang", "http://www.africau.edu/images/default/sample.pdf")
	if err != nil {
		log.Println("Info Error:", err.Error())
	}

	fmt.Println(status)
}
