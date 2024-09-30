package handlers

import (
	"log"
	"music-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Изменение данных песни
// @Description Обновляет инфорацию о песне по указанному ID
// @Tags Песни
// @Param id path int true "ID песни"
// @Param song body models.Song true "Обновленные данные о песне"
// @Success 200 {object} models.Song
// @Failure 400 {object} gin.H "Invalid ID or data"
// @Failure 404 {object} gin.H "Song not found"
// @Router /songs/{id} [put]

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
