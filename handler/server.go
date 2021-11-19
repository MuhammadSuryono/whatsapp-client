package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mri/whatsapp-client-message/models"
	"net/http"
)

func RunServer() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	r.NoMethod(func(context *gin.Context) {
		context.JSON(http.StatusMethodNotAllowed, models.CommonResponse{
			Code:      405,
			IsSuccess: false,
			Message:   "Method Not Allowed",
		})
		return
	})

	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, models.CommonResponse{
			Code:      404,
			IsSuccess: false,
			Message:   "Route Not Found",
		})
		return
	})

	return r
}
