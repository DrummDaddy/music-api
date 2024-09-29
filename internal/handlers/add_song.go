package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"music-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Songs []models.Song

// Метод для добавления новой песни
func AddSong(c *gin.Context) {
	var newSong models.Song
	if err := c.ShouldBindJSON(&newSong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}

	// Получение информации о песне из внешнего API
	apiURL := fmt.Sprintf("http://external-api-url/info?group=%s&song=%s", newSong.Artist, newSong.Title)
	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch song details from external API"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to read response from external API"})
		return
	}

	var songDetails struct {
		ReleaseDate string `json:"releaseDate"`
		Text        string `json:"text"`
		Link        string `json:"link"`
	}
	if err := json.Unmarshal(body, &songDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse external API response"})
		return
	}

	// Присваем новый ID для новой песни
	newSong.ID = uint(len(Songs) + 1)

	// Можно добавить дополнительные данные к песне
	newSong.Lyrics = append(newSong.Lyrics, songDetails.Text)

	// Добавляем новую песню в список
	Songs = append(Songs, newSong)
	c.JSON(http.StatusCreated, newSong)
}
