package main

import (
	"music-api/internal/config"
	"music-api/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация конфигурации
	config.LoadConfig()

	r := gin.Default()

	// Настройка маршрутов
	routers.SetupRoutes(r)

	// Запуск сервера
	r.Run(":8080")
}
