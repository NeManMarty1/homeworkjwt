package main

import (
	"project/internal/handlers"
	"project/internal/middleware"
	"project/internal/repository"
	"project/internal/services"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	userRepo := repository.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()

	authMiddleware := middleware.AuthMiddleware(utils.JWTSecret)

	r.POST("/login", userHandler.Login)
	r.POST("/register", userHandler.Register)

	protected := r.Group("/")
	protected.Use(authMiddleware)
	protected.GET("/profile", userHandler.GetProfile)

	r.Run(":8080")
}
