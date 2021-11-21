package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"mri/whatsapp-client-message/logs"
	"net/http"
	"os"
)

type IWhatsappClient interface {
	SendMessage(msidn string, message string) (bool, error)
	SendMessageWithDocument(msidn string, message string, urlFile string) (bool, error)
}

type WhatsappHandler struct {
	Token   string
	BaseUrl string
}

func NewWhatsappClientHandler() IWhatsappClient {
	return &WhatsappHandler{
		Token:   os.Getenv("NUSA_GATEWAY_TOKEN"),
		BaseUrl: os.Getenv("NUSA_GATEWAY_HOST"),
	}
}

func (wa *WhatsappHandler) SendMessage(msidn string, message string) (bool, error) {
	recLog := logs.NewLog()

	apiUrl := fmt.Sprintf("%s/send-message.php", wa.BaseUrl)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", wa.Token)
	_ = writer.WriteField("phone", msidn)
	_ = writer.WriteField("message", message)
	_ = writer.Close()

	recLog.WriteLog(recLog.MessageLogWithDate("Start request send message to " + msidn))
	recLog.WriteLog(recLog.MessageLogWithDate(message))

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, apiUrl, payload)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)
	buf, _ := ioutil.ReadAll(resp.Body)
	recLog.WriteLog(recLog.MessageLogWithDate("Resp send: " + string(buf)))
	if err != nil {
		recLog.WriteLog(recLog.MessageLogWithDate(fmt.Sprintf("Error send message: %v", err)))
		return false, err
	}

	return resp.StatusCode == 200, nil
}

func (wa *WhatsappHandler) SendMessageWithDocument(msidn string, message string, urlFile string) (bool, error) {
	recLog := logs.NewLog()

	apiUrl := fmt.Sprintf("%s/send-document.php", wa.BaseUrl)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", wa.Token)
	_ = writer.WriteField("phone", msidn)
	_ = writer.WriteField("document", urlFile)
	_ = writer.WriteField("caption", message)
	_ = writer.Close()

	recLog.WriteLog(recLog.MessageLogWithDate("Host sending: " + apiUrl))
	recLog.WriteLog(recLog.MessageLogWithDate("Start request send message to " + msidn))
	recLog.WriteLog(recLog.MessageLogWithDate(message))
	recLog.WriteLog(recLog.MessageLogWithDate(urlFile))

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, apiUrl, payload)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)
	buf, _ := ioutil.ReadAll(resp.Body)
	recLog.WriteLog(recLog.MessageLogWithDate("Resp send: " + string(buf)))
	if err != nil {
		recLog.WriteLog(recLog.MessageLogWithDate(fmt.Sprintf("Error send message: %v", err)))
		return false, err
	}

	return resp.StatusCode == 200, nil
}
