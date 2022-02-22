package service

import (
	"fmt"
	"log"
	"mri/whatsapp-client-message/db"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestSendMessageWA(t *testing.T) {
	errLoadEnv := godotenv.Load()
	db.InitConnectionFromEnvirontment().CreateNewConnection()
	//db.Connection.AutoMigrate(&models.LogWhatsapp{})
	if errLoadEnv != nil {
		log.Fatalln("Can't load env")
	}
	wa := NewWhatsappClientNusaGateWayHandler()
	provider := os.Getenv("PROVIDER")
	if provider == "OTHER" {
		apiClientWaOther := NewWhatsappClientHandler()
		_, err := apiClientWaOther.SendMessageOtherProvider("0895355698652", "Test from api golang wa sender")
		if err != nil {
			log.Println("Info Error:", err.Error())
		}
	} else {
		_, err := wa.SendMessage("0895355698652", "Test from api golang nusagateway")
		if err != nil {
			log.Println("Info Error:", err.Error())
		}
	}
}

func TestSendEmail(t *testing.T) {
	email := NewEmailHandler()
	param := ParamSendMessage{
		Recipients: "msuryono0@gmail.com",
		Subject:    "Notifikasi Error " + time.Now().Local().String(),
		Body:       "Tes kirim",
	}
	status, err := email.SendEmail(param)
	if err != nil {
		log.Println("Info Error:", err.Error())
	}

	fmt.Println(status)
}
