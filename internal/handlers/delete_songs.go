package handlers

import (
	"log"
	"music-api/internal/config"
	"music-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Удаление песни
// @Description Удаляет песню по указанному ID
// @Tags Песни
// @Param id path int true "ID песни"
// @Success 200 {object} gin.H "Song deleted"
// @Failure 400 {object} gin.H "Invalid ID"
// @Failure 500 {object} gin.H "Error deleting song"
// @Router /song/{id} [delete]

// Метод для удаления песни
func DeleteSong(c *gin.Context) {
	idParam := c.Param("id")
	log.Printf("DEBAG: Received request to delete sog with ID: %s", idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("INFO: Invalid ID provided: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	if err := config.DB.Delete(&models.Song{}, id).Error; err != nil {
		log.Printf("INFO: Error deleting song with ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting song"})
		return
	}

	log.Printf("INFO: Successfully deleted song with ID: %d", id)

	c.JSON(http.StatusOK, gin.H{"message": "Song deleted"})

}
