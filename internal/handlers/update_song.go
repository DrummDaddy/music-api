package handlers

import (
	"log"
	"music-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Метод для изменения данных песни
func UpdateSong(c *gin.Context) {
	log.Println("DEBUG: Get song by lirycs request receved")
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("INFO: Invalid ID provaided: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	var updatedSong models.Song
	if err := c.BindJSON(&updatedSong); err != nil {
		log.Printf("INFO: Invalid data provided: %s", err)
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
	log.Printf("INFO: Error, song not found: %d", id)
	c.JSON(http.StatusNotFound, gin.H{"message": "Song not found"})
}
