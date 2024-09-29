package routers

import (
	"music-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/songs", handlers.GetSongs)
	router.GET("/songs/:id/lyrics", handlers.GetLyrics)
	router.POST("/songs", handlers.AddSong)
	router.PUT("/songs/:id", handlers.UpdateSong)
	router.DELETE("/songs/:id", handlers.DeleteSong)
}
