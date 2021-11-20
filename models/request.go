package models

type ParamSendMessage struct {
	Msisdn  string `form:"msisdn" json:"msisdn"`
	Message string `form:"message" json:"message"`
}
