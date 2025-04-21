package postgres

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	localLog "bootstrap/internal/config/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresGORMConnection(ctx context.Context, databaseSource string) (*gorm.DB, error) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level (Silent, Error, Warn, Info)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error
			Colorful:                  true,        // Enable color
		},
	)

	localLog.Info(fmt.Sprintf("connection string: %v", databaseSource),
		zap.String("init", "NewPostgresGORMConnection"))

	db, err := gorm.Open(postgres.Open(databaseSource), &gorm.Config{Logger: newLogger})

	if err != nil {
		return nil, err
	}
	return db, nil
}
