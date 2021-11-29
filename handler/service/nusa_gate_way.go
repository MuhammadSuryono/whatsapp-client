package service

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"mri/whatsapp-client-message/logs"
	"net/http"
	"os"
)

type IWhatsappClientNusaGateWay interface {
	SendMessage(msidn string, message string) (bool, error)
	SendMessageWithDocument(msidn string, message string, urlFile string) (bool, error)
}

type NusaGateWayWhatsappHandler struct {
	Token   string
	BaseUrl string
}

func NewWhatsappClientNusaGateWayHandler() IWhatsappClientNusaGateWay {
	return &NusaGateWayWhatsappHandler{
		Token:   os.Getenv("NUSA_GATEWAY_TOKEN"),
		BaseUrl: os.Getenv("NUSA_GATEWAY_HOST"),
	}
}



func (wa *NusaGateWayWhatsappHandler) SendMessage(msidn string, message string) (bool, error) {
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
	if resp == nil {
		recLog.WriteLog(recLog.MessageLogWithDate(fmt.Sprintf("Null response: %v", err)))
		recLog.WriteToDbLog("NUSA_GATEWAY", msidn, message, "", 500, "Null response", fmt.Sprintf("Null response: %v", err))
		return false, errors.New("null response: " + fmt.Sprintf("%v", err))
		// return wa.SendMessageOtherProvider(msidn, message)
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	recLog.WriteLog(recLog.MessageLogWithDate("Resp send: " + string(buf)))
	if err != nil {
		recLog.WriteToDbLog("NUSA_GATEWAY", msidn, message, "", resp.StatusCode, string(buf), fmt.Sprintf("Error response: %v", err))
		recLog.WriteLog(recLog.MessageLogWithDate(fmt.Sprintf("Error send message: %v", err)))
		return false, err
		// return wa.SendMessageOtherProvider(msidn, message)
	}

	recLog.WriteToDbLog("NUSA_GATEWAY", msidn, message, "", resp.StatusCode, string(buf), fmt.Sprintf("Error response: %v", err))
	return resp.StatusCode == 200, nil
}

func (wa *NusaGateWayWhatsappHandler) SendMessageWithDocument(msidn string, message string, urlFile string) (bool, error) {
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
	if resp == nil {
		recLog.WriteLog(recLog.MessageLogWithDate(fmt.Sprintf("Null response: %v", err)))
		recLog.WriteToDbLog("NUSA_GATEWAY", msidn, message, urlFile, 500, "Null response", fmt.Sprintf("Null response: %v", err))
		return false, errors.New("null response: " + fmt.Sprintf("%v", err))
		// return wa.SendMessageOtherProvider(msidn, message)
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	recLog.WriteLog(recLog.MessageLogWithDate("Resp send: " + string(buf)))
	if err != nil {
		recLog.WriteToDbLog("NUSA_GATEWAY", msidn, message, urlFile, resp.StatusCode, string(buf), fmt.Sprintf("Error response: %v", err))
		recLog.WriteLog(recLog.MessageLogWithDate(fmt.Sprintf("Error send message: %v", err)))
		return false, err
		// return wa.SendMessageOtherProvider(msidn, message)
	}

	recLog.WriteToDbLog("NUSA_GATEWAY", msidn, message, urlFile, resp.StatusCode, string(buf), fmt.Sprintf("Error response: %v", err))
	return resp.StatusCode == 200, nil
}