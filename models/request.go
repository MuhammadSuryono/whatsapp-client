package models

type ParamSendMessage struct {
	Msisdn  string `form:"msisdn" json:"msisdn"`
	Message string `form:"message" json:"message"`
}

type ParamSendDocument struct {
	ParamSendMessage
	DocumentLink string `form:"document_link" json:"document_link"`
}
