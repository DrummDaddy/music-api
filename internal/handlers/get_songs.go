package handlers

import (
	"log"
	"music-api/internal/config"
	"music-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Получение списка песен
// @Description Возвращает список песен с возможностью фильтрации по исполнителю и альбому
// @Tags Песни
// @Param artist querry string false "Имя исполнителя"
// @Param album querry string false "Название альбома"
// @Param page querry int false "Номер страницы" default(1)
// @Param limit querry int false "Количество на странице" default(10)
// @Succes 200 {array} models.Song
// @Failure 500 {object} gin.H "Error retrieving songs"
// @Router /songs [get]

// Метод для получения данных библиотеки с фильтрацией и пагинацией
func GetSongs(c *gin.Context) {
	artistFilter := c.Query("artist")
	albumFilter := c.Query("album")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	log.Printf("DEBUG: Fetching songs with artist filter: %s and album filter: %s", artistFilter, albumFilter)

	query := config.DB
	if artistFilter != "" {
		query = query.Where("artist = ?", artistFilter)
	}
	if albumFilter != "" {
		query = query.Where("album = ?", albumFilter)
	}

	if err := query.Limit(limit).Offset((page - 1) * limit).Find(&Songs).Error; err != nil {
		log.Printf("INFO: Error retrieving songs from data base: %v", err)
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
	log.Printf("INFO: Total songs retrieved: %d", len(paginatedSongs))

	c.JSON(http.StatusOK, paginatedSongs)

}
