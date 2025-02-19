package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"service_notification/internal/config"
	_ "service_notification/internal/controller"
	controllers "service_notification/internal/controller"
	"service_notification/internal/email"
	migrations "service_notification/internal/migrations"
	"service_notification/internal/repository"
	services "service_notification/internal/service"
	"service_notification/internal/telegram"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	// Инициализация репозитория
	repo := repository.NewRepository(cfg)

	telega := telegram.NewTelegram(cfg)

	email := email.NewEmail(cfg)

	service := services.NewService(telega, email, repo)

	controller := controllers.NewNotificationController(service)

	migration := migrations.NewMigration(repo)

	ctx := context.Background()

	err = migration.InitTables(ctx, cfg)
	if err != nil {
		log.Fatal("Error creating table", err)
	}

	// Создание сервера
	server := gin.Default()

	// Маршруты
	server.POST("/notifications", controller.CreateNotificationHandler)

	// Запуск сервера
	if err := server.Run(":8080"); err != nil {
		log.Fatal("Server not started:", err)
	}
}
