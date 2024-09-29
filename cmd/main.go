package main

import (
	"music-api/internal/config"
	"music-api/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация конфигурации
	config.LoadConfig()
	config.ConnectDatabase()

	r := gin.Default()
	routers.SetupRoutes(r)

	// Настройка маршрутов
	routers.SetupRoutes(r)

	// Запуск сервера
	r.Run(":8080")
}
