package main

import (
	"log"
	_ "music-api/docs"
	"music-api/internal/config"
	"music-api/internal/handlers"
	"music-api/internal/logger"
	"music-api/internal/routers"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"
)

func main() {

	// Инициализация логгера
	myLogger, err := logger.NewLogger(logger.INFO, "app.log")
	if err != nil {
		log.Fatalf("Ошибка создания логгера: %v", err)
	}
	defer myLogger.Close()

	// Пример инициализации обработчиков с использованием логгера
	handlers.InitializeLogger(myLogger)

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
