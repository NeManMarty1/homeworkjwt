package app

import (
	"context"
	"homeworkjwt/internal/config"
)

// Run запускает основной процесс сервера, включая инициализацию всех зависимостей
func Run(ctx context.Context, cfg *config.Config) error {
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

	// TODO
	// gracefull shutdown

	r.Run(fmt.Sprintf(":%s", cfg.HTTP.Port))
}