package handlers

import (
	"music-api/internal/models"

	"github.com/gin-gonic/gin"
)

// @Summary Получить все песни
// @Description Получить список всех песен
// @Produce json
// @Success 200 {array} models.Song
// @Router /songs [get]

func GetAllSongs(c *gin.Context) {
	songs := []models.Song{
		{ID: 1, Title: "Song 1", Artist: "Artist 1"},
		{ID: 2, Title: "Song 2", Artist: "Artist 2"},
	}
	c.JSON(200, songs)
}
