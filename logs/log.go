package logs

import (
	"fmt"
	"log"
	"os"
	"time"
)

type ILog interface {
	WriteLog(messageLog interface{})
	MessageLogWithDate(message string) string
}

type LogHandler struct {
}

func NewLog() ILog {
	return &LogHandler{}
}

func (l *LogHandler) WriteLog(messageLog interface{}) {
	filename := fmt.Sprintf("%s-%d-%02d-%02d", "../logs/logs/log", time.Now().Year(), time.Now().Month(), time.Now().Day())
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
