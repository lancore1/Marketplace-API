package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	// this will automatically load your .env file: (blank-import / side-effect import)
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	DB   PostgresConfig
	Port string
}

type PostgresConfig struct {
	Username string
	Password string
	URL      string
	Port     string
}

func hasEmptyFields(s interface{}) error {
	value := reflect.ValueOf(s)

	// Convert to value if we get ptr
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// Check Struct only
	if value.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		if field.Kind() == reflect.Struct {
			if err := hasEmptyFields(field.Interface()); err != nil {
				return err
			}
		}

		// Check String
		if field.Kind() == reflect.String {
			if strings.TrimSpace(field.String()) == "" {
				return fmt.Errorf("field '%s' is empty", value.Type().Field(i).Name)
			}
		}
	}
	return nil
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port: os.Getenv("PORT"),
		DB: PostgresConfig{
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PWD"),
			URL:      os.Getenv("POSTGRES_URL"),
			Port:     os.Getenv("POSTGRES_PORT"),
		},
	}

	err := hasEmptyFields(cfg)

	if err != nil {
		return nil, fmt.Errorf("unable to load config: %w", err)
	}

	return cfg, nil
}
