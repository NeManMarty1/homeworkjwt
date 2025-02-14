package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"homeworkjwt/internal/config"
	"homeworkjwt/internal/handlers"
	"homeworkjwt/internal/middleware"
	"homeworkjwt/internal/repository"
	"homeworkjwt/internal/services"
)

func main() {
	// Получение конфигурации приложения
	cfg, err := config.GetConfigFromEnv()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s\n", err.Error())
	}

	userRepo := repository.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()

	authMiddleware := middleware.AuthMiddleware(cfg.JWT.Secret)

	r.POST("/login", userHandler.Login)
	r.POST("/register", userHandler.Register)

	protected := r.Group("/")
	protected.Use(authMiddleware)
	protected.GET("/profile", userHandler.GetProfile)

	r.Run(fmt.Sprintf(":%s", cfg.HTTP.Port))
}
