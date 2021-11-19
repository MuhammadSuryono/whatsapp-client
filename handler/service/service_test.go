package service

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestSendMessageWA(t *testing.T) {
	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil {
		log.Fatalln("Can't load env")
	}
	wa := NewWhatsappClientHandler(os.Getenv("WA_SENDER_HOST"), os.Getenv("WA_SENDER_TOKEN"))
	status, err := wa.SendMessage("082186826264", "Test from api golang")
	if err != nil {
		log.Println("Info Error:", err.Error())
	}

	fmt.Println(status)
}