package routers

import (
	"music-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/songs", handlers.GetSongs)
	router.GET("/songs/:id", handlers.GetSongByID)

}
