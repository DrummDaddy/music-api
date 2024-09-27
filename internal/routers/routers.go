package routers

import (
	"music-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("songs", handlers.GetAllSongs)
	}

}
