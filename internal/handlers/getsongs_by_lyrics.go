package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Метод для получения текста песни с пагинацией по куплетам
func GetLyrics(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	for _, song := range Songs {
		if song.ID == uint(id) {
			page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
			limit, _ := strconv.Atoi(c.DefaultQuery("limit", "1"))

			start := (page - 1) * limit
			end := start + limit
			if start > len(song.Lyrics) {
				start = len(song.Lyrics)
			}
			paginatedLyrics := song.Lyrics[start:end]

			c.JSON(http.StatusOK, paginatedLyrics)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Song not found"})
}
