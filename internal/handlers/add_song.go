package handlers

import (
	"music-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//var Songs []models.Song

// Метод для добавления новой песни

func AddSong(c *gin.Context) {
	var newSong models.Song
	if err := c.ShouldBindJSON(&newSong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}

	// Присваем новый ID для новой песни
	newSong.ID = uint(len(Songs) + 1)

	// Добавляем новую песню в список
	Songs = append(Songs, newSong)
	c.JSON(http.StatusCreated, newSong)

}
