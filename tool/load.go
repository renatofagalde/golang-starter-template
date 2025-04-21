package tool

import (
	"fmt"
	"os"
	"time"
)

type ToolLoadEnvironmet struct{}

type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func (le *ToolLoadEnvironmet) Do() (config Config, err error) {

	config.DBDriver = os.Getenv("DB_DRIVER")
	config.DBSource = os.Getenv("DB_SOURCE")
	config.ServerAddress = os.Getenv("SERVER_ADDRESS")
	config.TokenSymmetricKey = os.Getenv("TOKEN_SYMMETRIC_KEY")

	accessTokenDurationStr := os.Getenv("ACCESS_TOKEN_DURATION")
	if accessTokenDurationStr == "" {
		err = fmt.Errorf("variável de ambiente ACCESS_TOKEN_DURATION não está definida ou está vazia")
		return
	}

	config.AccessTokenDuration, err = time.ParseDuration(accessTokenDurationStr)
	if err != nil {
		err = fmt.Errorf("erro ao converter ACCESS_TOKEN_DURATION para time.Duration: %w", err)
		return
	}

	if config.DBDriver == "" || config.ServerAddress == "" || config.TokenSymmetricKey == "" {
		err = fmt.Errorf("algumas variáveis obrigatórias não estão definidas no ambiente")
		return
	}

	config.DBSource = le.buildDBSource(config.DBSource)
	return
}

func (le *ToolLoadEnvironmet) buildDBSource(defaultDBSource string) string {
	dbUser := os.Getenv("app_database_user")
	dbPassword := os.Getenv("app_database_password")
	dbURL := os.Getenv("app_database_url")

	if dbUser != "" && dbPassword != "" && dbURL != "" {
		return fmt.Sprintf("postgresql://%s:%s@%s", dbUser, dbPassword, dbURL)
	}
	return defaultDBSource
}
