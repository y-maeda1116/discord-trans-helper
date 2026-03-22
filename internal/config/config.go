package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	DiscordToken string
	DeepLAuthKey string
}

// Load reads environment variables from .env file and returns Config
func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := &Config{
		DiscordToken: os.Getenv("DISCORD_TOKEN"),
		DeepLAuthKey: os.Getenv("DEEPL_AUTH_KEY"),
	}

	if cfg.DiscordToken == "" {
		log.Fatal("DISCORD_TOKEN is required")
	}
	if cfg.DeepLAuthKey == "" {
		log.Fatal("DEEPL_AUTH_KEY is required")
	}

	return cfg
}
