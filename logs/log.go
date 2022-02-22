package logs

import (
	"fmt"
	"log"
	"mri/whatsapp-client-message/db"
	"mri/whatsapp-client-message/models"
	"os"
	"time"
)

type ILog interface {
	WriteLog(messageLog interface{})
	MessageLogWithDate(message string) string
	WriteToDbLog(provider, to, message, documentLink string, statusCode int, response, error string, res models.ResponseWhatsappProvider)
	WriteToDbLogResend(idMessage int64, status bool)
}

type LogHandler struct {
}

func NewLog() ILog {
	return &LogHandler{}
}

func (l *LogHandler) WriteLog(messageLog interface{}) {
	filename := fmt.Sprintf("%s-%d-%02d-%02d", "../../logs/logs/log", time.Now().Year(), time.Now().Month(), time.Now().Day())
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Sprintf("error opening file: %v", err)
		return
	}
	defer f.Close()
	message := fmt.Sprintf("> %v", messageLog)
	log.SetOutput(f)
	log.Println(message)
}

func (l *LogHandler) MessageLogWithDate(message string) string {
	return fmt.Sprintf("%d-%02d-%02d -> %s", time.Now().Year(), time.Now().Month(), time.Now().Day(), message)
}

func (l *LogHandler) WriteToDbLog(provider, to, message, documentLink string, statusCode int, response, error string, res models.ResponseWhatsappProvider) {
	status := true
	if statusCode >= 400 {
		status = false
	}

	logs := models.LogWhatsapp{
		ApplicationCode: "budget-001",
		Provider:        provider,
		To:              to,
		Message:         message,
		IdMessage:       res.Id,
		StatusSend:      res.Status,
		Response:        response,
		Errors:          error,
		StatusCode:      statusCode,
		Status:          status,
		DocumentLink:    documentLink,
	}

	db.Connection.Create(&logs)
}

func (l *LogHandler) WriteToDbLogResend(idMessage int64, status bool) {
	logResend := getLogResend(idMessage, status)
	logResend.CountResend += 1
	if status {
		logResend.Status = true
		logResend.StatusCode = 200
	}
	db.Connection.Save(&logResend)

}

func getLogResend(idMessage int64, status bool) (log models.LogWhatsapp) {
	db.Connection.Where("id = ? AND count_resend <= ?", idMessage, 3).First(&log)
	return
}
