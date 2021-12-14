package models

import "time"

type CommonResponse struct {
	Code      int         `json:"code"`
	IsSuccess bool        `json:"is_success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

type LogWhatsapp struct {
	Id              int64     `json:"id" gorm:"primaryKey"`
	ApplicationCode string    `json:"application_code"`
	Provider        string    `json:"provider"`
	To              string    `json:"to"`
	Message         string    `json:"message" gorm:"TEXT"`
	DocumentLink    string    `json:"document_link" gorm:"TEXT"`
	Response        string    `json:"response" gorm:"TEXT"`
	Errors          string    `json:"errors" gorm:"TEXT"`
	StatusCode      int       `json:"status_code"`
	Status          bool      `json:"status"`
	CountResend		int	  	  `json:"count_resend" gorm:"DEFAULT:0"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Config struct {
	Id              int64     `json:"id" gorm:"primaryKey"`
	ApplicationCode string    `json:"application_code"`
	Services		string	  `json:"services"`
	NeedResend 		bool 	  `json:"need_resend" gorm:"DEFAULT:0"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
