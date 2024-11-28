package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"music-api/internal/config"
	"music-api/internal/logger"

	"music-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Добавление новой песни
// @Description Добавляет новую песню в библиотеку
// @Tags Песни
// @Param song body models.Song true "Данные новой песни"
// @Success 201 {object} models.Song
// @Failure 400 {object} gin.H "Invalid JSON data"
// @Failure 500 {object} gin.H "Could not add song data base"s
// @Router /songs [post]

// InitializeLogger инициализирует кастомный логгер для пакета
func InitializeLogger(customLogger *logger.Logger) {
	appLog = customLogger
}

var Songs []models.Song

// Метод для добавления новой песни
func AddSong(c *gin.Context) {
	log.Println("DEBUG: Addsong request receved")
	var newSong models.Song
	if err := c.ShouldBindJSON(&newSong); err != nil {
		log.Printf("INFO: Invalid JSON data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}

	if err := config.DB.Create(&newSong).Error; err != nil {
		log.Printf("INFO: Could not add song to database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not add song to database"})
		return

	}

	// Получение информации о песне из внешнего API
	apiURL := fmt.Sprintf("http://external-api-url/info?group=%s&song=%s", newSong.Artist, newSong.Title)
	log.Printf("DEBUG: Calling external API: %s", apiURL)
	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("INFO: Failed to fetch song details from external API: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch song details from external API"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("INFO: Failed to read response from external API: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to read response from external API"})
		return
	}

	var songDetails struct {
		ReleaseDate string `json:"releaseDate"`
		Text        string `json:"text"`
		Link        string `json:"link"`
	}
	if err := json.Unmarshal(body, &songDetails); err != nil {
		log.Printf("INFO: Failed to parce external API response: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse external API response"})
		return
	}

	log.Printf("INFO: External API call successful, details fetched")

	// Присваем новый ID для новой песни
	newSong.ID = uint(len(Songs) + 1)

	// Можно добавить дополнительные данные к песне
	newSong.Lyrics = append(newSong.Lyrics, songDetails.Text)

	// Добавляем новую песню в список
	Songs = append(Songs, newSong)
	c.JSON(http.StatusCreated, newSong)
}
