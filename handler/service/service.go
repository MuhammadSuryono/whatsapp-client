package service

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"mri/whatsapp-client-message/logs"
	"net/http"
	"net/url"
	"os"
	"strings"
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
	if resp == nil {
		recLog.WriteLog(recLog.MessageLogWithDate(fmt.Sprintf("Null response: %v", err)))
		recLog.WriteToDbLog("NUSA_GATEWAY", msidn, message, "", 500, "Null response", fmt.Sprintf("Null response: %v", err))
		//return false, errors.New("null response: " + fmt.Sprintf("%v", err))
		return wa.SendMessageOtherProvider(msidn, message)
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	recLog.WriteLog(recLog.MessageLogWithDate("Resp send: " + string(buf)))
	if err != nil {
		recLog.WriteToDbLog("NUSA_GATEWAY", msidn, message, "", resp.StatusCode, string(buf), fmt.Sprintf("Error response: %v", err))
		recLog.WriteLog(recLog.MessageLogWithDate(fmt.Sprintf("Error send message: %v", err)))
		//return false, err
		return wa.SendMessageOtherProvider(msidn, message)
	}

	recLog.WriteToDbLog("NUSA_GATEWAY", msidn, message, "", resp.StatusCode, string(buf), fmt.Sprintf("Error response: %v", err))
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
	if resp == nil {
		recLog.WriteLog(recLog.MessageLogWithDate(fmt.Sprintf("Null response: %v", err)))
		recLog.WriteToDbLog("NUSA_GATEWAY", msidn, message, urlFile, 500, "Null response", fmt.Sprintf("Null response: %v", err))
		//return false, errors.New("null response: " + fmt.Sprintf("%v", err))
		return wa.SendMessageOtherProvider(msidn, message)
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	recLog.WriteLog(recLog.MessageLogWithDate("Resp send: " + string(buf)))
	if err != nil {
		recLog.WriteToDbLog("NUSA_GATEWAY", msidn, message, urlFile, resp.StatusCode, string(buf), fmt.Sprintf("Error response: %v", err))
		recLog.WriteLog(recLog.MessageLogWithDate(fmt.Sprintf("Error send message: %v", err)))
		//return false, err
		return wa.SendMessageOtherProvider(msidn, message)
	}

	recLog.WriteToDbLog("NUSA_GATEWAY", msidn, message, urlFile, resp.StatusCode, string(buf), fmt.Sprintf("Error response: %v", err))
	return resp.StatusCode == 200, nil
}

func (wa *WhatsappHandler) SendMessageOtherProvider(msidn string, message string) (bool, error) {
	recLog := logs.NewLog()
	wa.Token = os.Getenv("WA_SENDER_TOKEN")
	wa.BaseUrl = os.Getenv("WA_SENDER_HOST")

	apiUrl := fmt.Sprintf("%s/kirim_wa", wa.BaseUrl)
	dataPost := url.Values{}
	dataPost.Set("no_wa", msidn)
	dataPost.Set("pesan", message)

	recLog.WriteLog(recLog.MessageLogWithDate("Start request send message to " + msidn))
	recLog.WriteLog(recLog.MessageLogWithDate(message))

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, apiUrl, strings.NewReader(dataPost.Encode()))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+wa.Token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if resp == nil {
		recLog.WriteLog(recLog.MessageLogWithDate(fmt.Sprintf("Null response: %v", err)))
		recLog.WriteToDbLog("WA_SENDER", msidn, message, "", 500, "Null response", fmt.Sprintf("Null response: %v", err))
		return false, errors.New("null response: " + fmt.Sprintf("%v", err))
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	recLog.WriteLog(recLog.MessageLogWithDate("Resp send: " + string(buf)))
	if err != nil {
		recLog.WriteToDbLog("WA_SENDER", msidn, message, "", resp.StatusCode, string(buf), fmt.Sprintf("Error response: %v", err))
		recLog.WriteLog(recLog.MessageLogWithDate(fmt.Sprintf("Error send message: %v", err)))
		return false, err
	}

	recLog.WriteToDbLog("WA_SENDER", msidn, message, "", resp.StatusCode, string(buf), fmt.Sprintf("Error response: %v", err))
	return resp.StatusCode == 200, nil
}
