package handlers

import (
	"music-api/internal/config"
	"music-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Метод для удаления песни
func DeleteSong(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	if err := config.DB.Delete(&models.Song{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting song"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song deleted"})

}
