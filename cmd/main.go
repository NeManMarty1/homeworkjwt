package main

import (
	"homeworkjwt/internal/config"
	"homeworkjwt/internal/handlers"
	"homeworkjwt/internal/middleware"
	"homeworkjwt/internal/repository"
	"homeworkjwt/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()

	repo := repository.NewUserRepository()
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		token, err := service.Login(input.Email, input.Password)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"token": token})
	})

	r.GET("/profile", middleware.AuthMiddleware(cfg.JWTSecret), handler.GetProfile)

	r.Run(":" + cfg.AppPort)
}
