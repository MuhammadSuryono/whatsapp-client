package service

import (
	"fmt"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestSendMessageWA(t *testing.T) {
	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil {
		log.Fatalln("Can't load env")
	}
	wa := NewWhatsappClientHandler()
	status, err := wa.SendMessageWithDocument("0895355698652", "Test from api golang", "http://180.211.92.131/dev-budget/document/yyuwQ73pQXEg2LF.pdf")
	if err != nil {
		log.Println("Info Error:", err.Error())
	}

	fmt.Println(status)
}
