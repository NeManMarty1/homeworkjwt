package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"homeworkjwt/internal/config"
	"homeworkjwt/internal/handlers"
	"homeworkjwt/internal/middleware"

	migrate "homeworkjwt/internal/app"
	"homeworkjwt/internal/pgdb"
	"homeworkjwt/internal/postgres"
	"homeworkjwt/internal/services"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.GetConfigFromEnv()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s\n", err.Error())
	}

	// TODO
	// Инициализация интанса логера 

	if err = app.Run(ctx, cfg)
}
