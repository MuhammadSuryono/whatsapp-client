package service

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

const (
	EmailUrl = "http://192.168.8.2:8081/api/v1/email"
)

type ParamSendMessage struct {
	Recipients   string `form:"recipients" json:"recipients"`
	RecipientsCC string `form:"recipients_cc" json:"recipients_cc"`
	Subject      string `form:"subject" json:"subject"`
	TypeBody     string `form:"type_body" json:"type_body"`
	Body         string `form:"body" json:"body"`
}

type IEmailService interface {
	SendEmail(param ParamSendMessage) (bool, error)
}

type EmailHandler struct {
}

func NewEmailHandler() IEmailService {
	return &EmailHandler{}
}

func (em *EmailHandler) SendEmail(param ParamSendMessage) (bool, error) {
	apiUrl := fmt.Sprintf("%s/send-notification-message", EmailUrl)
	fmt.Println(apiUrl)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("recipients", param.Recipients)
	_ = writer.WriteField("recipients_cc", param.RecipientsCC)
	_ = writer.WriteField("subject", param.Subject)
	_ = writer.WriteField("type_body", param.TypeBody)
	_ = writer.WriteField("body", param.Body)
	_ = writer.Close()

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, apiUrl, payload)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)
	if resp == nil {
		return false, errors.New("null response: " + fmt.Sprintf("%v", err))
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Resp : " + string(buf))
	if err != nil {
		return false, err
	}

	return resp.StatusCode == 200, nil
}
