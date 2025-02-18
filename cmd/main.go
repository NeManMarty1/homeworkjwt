package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"homeworkjwt/internal/config"
	"homeworkjwt/internal/handlers"
	"homeworkjwt/internal/middleware"

	// "homeworkjwt/internal/repository"
	migrate "homeworkjwt/internal/app"
	"homeworkjwt/internal/pgdb"
	"homeworkjwt/internal/postgres"
	"homeworkjwt/internal/services"
)

func main() {
	// Получение конфигурации приложения
	cfg, err := config.GetConfigFromEnv()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s\n", err.Error())
	}

	// Инициализация клиента с PostgreSQL
	pool := postgres.New(cfg)

	err = migrate.InitMigrations()
	if err != nil {
		log.Fatalf("Failed to initialize migrations: %s", err.Error())
	}
	// Инициализация репоизториев, используя пул соединений с PostgreSQL
	repositories := pgdb.NewRepositries(pool)

	// userRepo := repository.NewUserRepository()
	userService := services.NewUserService(repositories)
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
