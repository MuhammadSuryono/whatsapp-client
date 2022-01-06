package models

import (
	"mri/whatsapp-client-message/db"
	"time"
)

func GetLogWhatsapp() (logs []LogWhatsapp){
	_= db.Connection.Limit(10).Order("id desc").Find(&logs)
	return
}

func TotalLogs() (total int64) {
	_= db.Connection.Model(&LogWhatsapp{}).Count(&total)
	return
}

func LastUpdateLogs() time.Time {
	var log LogWhatsapp
	_= db.Connection.Last(&log)
	return log.UpdatedAt
}