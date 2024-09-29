package handlers

import (
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

	for i, song := range Songs {
		if song.ID == uint(id) {
			Songs = append(Songs[:i], Songs[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Song deleted"})
			return
		}

	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Song not found"})
}
