package models

type ParamSendMessage struct {
	Msisdn string `json:"msisdn" binding:"required"`
	Message string `json:"message" binding:"required"`
}
