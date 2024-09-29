package handlers

import (
	"music-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Метод для изменения данных песни
func UpdateSong(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	var updatedSong models.Song
	if err := c.BindJSON(&updatedSong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	for i, song := range Songs {
		if song.ID == uint(id) {
			updatedSong.ID = song.ID
			Songs[i] = updatedSong
			c.JSON(http.StatusOK, Songs[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Song not found"})
}
