package models

import (
	"mri/whatsapp-client-message/db"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetLogWhatsapp(c *gin.Context) (logs []LogWhatsapp) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	offset := 0

	if page == 0 {
		page = 1
	}

	if size == 0 {
		size = 10
	}

	if page > 1 {
		offset = (page - 1) * size
	}
	_ = db.Connection.Limit(size).Offset(offset).Order("id desc").Find(&logs)
	return
}

func TotalLogs() (total int64) {
	_ = db.Connection.Model(&LogWhatsapp{}).Count(&total)
	return
}

func LastUpdateLogs() time.Time {
	var log LogWhatsapp
	_ = db.Connection.Last(&log)
	return log.UpdatedAt
}
