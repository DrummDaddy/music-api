package handlers

import (
	"music-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Songs хранит список песен
var Songs = []models.Song{
	{ID: 1, Title: "Song One", Artist: "Artist One", Album: "Album One"},
	{ID: 2, Title: "Song Two", Artist: "Artist Two", Album: "Album Two"},
}

// GetSongs возвращает список всех песен
func GetSongs(c *gin.Context) {
	c.JSON(http.StatusOK, Songs)
}

// GetSongByID возвращает песню по ID
func GetSongByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	for _, song := range Songs {
		if song.ID == uint(id) {
			c.JSON(http.StatusOK, song)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Song not found"})
}
