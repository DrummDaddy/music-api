package handlers

import (
	"music-api/internal/config"
	"music-api/internal/logger"
	"music-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// log – глобальная переменная для логирования в этом пакете.
var appLog *logger.Logger

// InitLogger инициализирует логгер для использования в этом пакете.
func InitLogger(customLogger *logger.Logger) {
	appLog = customLogger
}

// GetSongs обрабатывает HTTP-запросы на получение песен.
// Он поддерживает фильтрацию по артисту и альбому, а также пагинацию.
func GetSongs(c *gin.Context) {

	artistFilter := c.Query("artist")
	albumFilter := c.Query("album")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Логируем информацию о начале процесса получения песен с параметрами фильтрации.
	appLog.DEBUG("Fetching songs with artist filter: ", artistFilter, " and album filter: ", albumFilter)

	// Инициализируем запрос к базе данных, предварительно загружая связанных артистов.
	query := config.DB.Preload("Artist")

	if artistFilter != "" {
		// Присоединяем таблицу артистов и фильтруем по имени артиста.
		query = query.Joins("JOIN artists ON artists.id = songs.artist_id").Where("artists.name = ?", artistFilter)
	}
	if albumFilter != "" {
		// Фильтруем по названию альбома.
		query = query.Where("album = ?", albumFilter)
	}

	// Создаём слайс для хранения песен.
	var Songs []models.Song

	// Выполняем запрос к базе данных с заданными ограничениям и смещением (пагинация).
	if err := query.Limit(limit).Offset((page - 1) * limit).Find(&Songs).Error; err != nil {
		appLog.Error("Error retrieving songs from database: ", err)
		// Если возникла ошибка, возвращаем статус 500 и сообщение об ошибке.
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving songs"})
		return
	}

	// Создаём слайс для хранения отфильтрованных песен.
	var filteredSongs []models.Song

	// Фильтрация песен по названию артиста и албома.
	for _, song := range Songs {
		if (artistFilter == "" || song.Artist.Name == artistFilter) && (albumFilter == "" || song.Album == albumFilter) {
			filteredSongs = append(filteredSongs, song)
		}
	}

	// Рассчитываем начальный и конечный индексы для пагинации отфильтрованных песен.
	start := (page - 1) * limit
	end := start + limit
	if start > len(filteredSongs) {
		start = len(filteredSongs)
	}
	if end > len(filteredSongs) {
		end = len(filteredSongs)
	}

	// Получаем подмножество песен после фильтрации и пагинации.
	paginatedSongs := filteredSongs[start:end]
	appLog.Info("Total songs retrieved: ", len(paginatedSongs))

	// Возвращаем отфильтрованные и пагинированные песни в JSON-формате.
	c.JSON(http.StatusOK, paginatedSongs)
}
