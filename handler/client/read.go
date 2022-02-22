package client

import (
	"mri/whatsapp-client-message/models"

	"github.com/gin-gonic/gin"
)

func (cl *ClientHandlerWhatsapp) GetLogWhatsapp(c *gin.Context) {
	logWhatsapps := models.GetLogWhatsapp(c)
	totalLogs := models.TotalLogs()
	lastUpdate := models.LastUpdateLogs()

	c.JSON(200, models.CommonResponse{
		Code:      200,
		IsSuccess: true,
		Message:   "Success retrieve data",
		Data: models.ResponseListLog{
			TotalData:  totalLogs,
			Records:    logWhatsapps,
			LastUpdate: lastUpdate,
		},
	})
}
