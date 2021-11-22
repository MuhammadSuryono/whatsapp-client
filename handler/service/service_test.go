package service

import (
	"fmt"
	"log"
	"mri/whatsapp-client-message/db"
	"testing"

	"github.com/joho/godotenv"
)

func TestSendMessageWA(t *testing.T) {
	errLoadEnv := godotenv.Load()
	db.InitConnectionFromEnvirontment().CreateNewConnection()
	if errLoadEnv != nil {
		log.Fatalln("Can't load env")
	}
	wa := NewWhatsappClientHandler()
	status, err := wa.SendMessage("0895355698652", "Test from api golang")
	if err != nil {
		log.Println("Info Error:", err.Error())
	}

	fmt.Println(status)
}
