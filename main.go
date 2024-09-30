package main

import (
	_ "music-api/docs"
	"music-api/internal/config"
	"music-api/internal/routers"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"
)

func main() {
	// Инициализация конфигурации
	config.LoadConfig()
	config.ConnectDatabase()

	r := gin.Default()
	routers.SetupRoutes(r)

	// Настройка маршрутов
	routers.SetupRoutes(r)

	// подключение Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера
	r.Run(":8080")
}
