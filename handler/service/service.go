package service

import (
	"fmt"
	"io/ioutil"
	"mri/whatsapp-client-message/logs"
	"net/http"
	"net/url"
	"strings"
)

type IWhatsappClient interface {
	SendMessage(msidn string, message string) (bool, error)
}

type WhatsappHandler struct {
	Token   string
	BaseUrl string
}

func NewWhatsappClientHandler(baseUrl, token string) IWhatsappClient {
	return &WhatsappHandler{
		Token:   token,
		BaseUrl: baseUrl,
	}
}

func (wa *WhatsappHandler) SendMessage(msidn string, message string) (bool, error) {
	recLog := logs.NewLog()

	apiUrl := fmt.Sprintf("%s/api/kirim_wa", wa.BaseUrl)
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
	recLog.WriteLog(recLog.MessageLogWithDate("Err send: " + err.Error()))
	buf, _ := ioutil.ReadAll(resp.Body)
	recLog.WriteLog(recLog.MessageLogWithDate("Resp send: " + string(buf)))
	if err != nil {
		recLog.WriteLog(recLog.MessageLogWithDate("Error send message: " + err.Error()))
		return false, err
	}

	return resp.StatusCode == 200, nil
}
