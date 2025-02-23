package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	DbName         string `env:"DB_NAME"`
	DbUser         string `env:"DB_USER"`
	DbPassword     string `env:"DB_PASSWORD"`
	DbHost         string `env:"DB_HOST"`
	DbPort         int    `env:"DB_PORT"`
	TelegramToken  string `env:"TELEGRAM_TOKEN"`
	SMTPEmail      string `env:"SMTP_EMAIL"`
	SMTPPassword   string `env:"SMTP_PASSWORD"`
	SMTPHost       string `env:"SMTP_HOST"`
	SMTPPort       int    `env:"SMTP_PORT"`
	TestUserMail   string `env:"TEST_USER_EMAIL"`
	TelegramChatID string `env:"TELEGRAM_CHAT_ID"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg.DbName = os.Getenv("DB_NAME")
	cfg.DbUser = os.Getenv("DB_USER")
	cfg.DbPassword = os.Getenv("DB_PASSWORD")
	cfg.DbHost = os.Getenv("DB_HOST")
	cfg.DbPort, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	cfg.TelegramToken = os.Getenv("TELEGRAM_TOKEN")
	cfg.SMTPEmail = os.Getenv("SMTP_EMAIL")
	cfg.SMTPPassword = os.Getenv("SMTP_PASSWORD")
	cfg.SMTPHost = os.Getenv("SMTP_HOST")
	cfg.SMTPPort, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
	cfg.TestUserMail = os.Getenv("TEST_USER_EMAIL")
	cfg.TelegramChatID = os.Getenv("TELEGRAM_CHAT_ID")

	return &cfg, nil
}
