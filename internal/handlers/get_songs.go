package handlers

import (
	"music-api/internal/config"
	"music-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Метод для получения данных библиотеки с фильтрацией и пагинацией
func GetSongs(c *gin.Context) {
	artistFilter := c.Query("artist")
	albumFilter := c.Query("album")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	query := config.DB
	if artistFilter != "" {
		query = query.Where("artist = ?", artistFilter)
	}
	if albumFilter != "" {
		query = query.Where("album = ?", albumFilter)
	}

	if err := query.Limit(limit).Offset((page - 1) * limit).Find(&Songs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retreving songs"})
		return
	}

	c.JSON(http.StatusOK, Songs)

	var filteredSongs []models.Song

	//Фильтрация песен
	for _, song := range Songs {
		if (artistFilter == "" || song.Artist == artistFilter) && (albumFilter == "" || song.Album == albumFilter) {
			filteredSongs = append(filteredSongs, song)
		}
	}

	// Пагинация
	start := (page - 1) * limit
	end := start + limit
	if start > len(filteredSongs) {
		start = len(filteredSongs)
	}
	if end > len(filteredSongs) {
		end = len(filteredSongs)
	}
	paginatedSongs := filteredSongs[start:end]

	c.JSON(http.StatusOK, paginatedSongs)

}
