package migrate

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	defaultAttempts = 5
	defaultTimeout  = time.Second
)

// InitMigrations запускает миграцию, принимая строку подключения к базе данных
func InitMigrations() error {
	// Получаем строку подключения к базе данных из переменной окружения
	databaseURL, ok := os.LookupEnv("PG_URL")
	if !ok || len(databaseURL) == 0 {
		log.Fatalf("migrate: environment variable not declared: PG_URL")
	}

	var (
		attempts = defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	// Цикл попыток подключения
	for attempts > 0 {
		m, err = migrate.New("file://migrations", databaseURL)
		if err == nil {
			break
		}

		log.Printf("Migrate: pgdb is trying to connect, attempts left: %d", attempts)
		time.Sleep(defaultTimeout)
		attempts--
	}

	// Проверка на окончание попыток и успешное подключение
	if err != nil {
		log.Fatalf("Migrate: pgdb connect error: %s", err)
		return err
	}

	// Выполнение миграции
	err = m.Up()
	defer func() { _, _ = m.Close() }()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: up error: %s", err)
		return err
	}

	// Если миграции нет (состояние без изменений)
	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Migrate: no change")
		return nil
	}

	log.Printf("Migrate: up success")

	return nil
}
