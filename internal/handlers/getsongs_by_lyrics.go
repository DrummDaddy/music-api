package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Метод для получения текста песни с пагинацией по куплетам
func GetLyrics(c *gin.Context) {
	log.Println("DEBUG: Get song by lirycs request receved")
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("INFO: Invalid ID provaided: %s", idParam)
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
	log.Printf("INFO: Error song not found, %d", id)
	c.JSON(http.StatusNotFound, gin.H{"message": "Song not found"})
}
