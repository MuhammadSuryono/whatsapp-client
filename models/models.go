package models

type CommonResponse struct {
	Code       int         `json:"code"`
	IsSuccess  bool        `json:"is_success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

