package configs

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

type Config struct {
	Server *ServerConfig `mapstructure:"http_server"`
	DB     *DBConfig     `mapstructure:"db"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Name     string `mapstructure:"name"`
	Password string
}

func Load() *Config {
	config := &Config{
		Server: &ServerConfig{},
		DB:     &DBConfig{},
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msgf("Error loading .env file")
	}

	config.Server.Host = os.Getenv("SERVER_HOST")
	config.Server.Port = os.Getenv("SERVER_PORT")

	config.DB.Host = os.Getenv("DB_HOST")
	config.DB.Port = os.Getenv("DB_PORT")
	config.DB.User = os.Getenv("DB_USER")
	config.DB.Name = os.Getenv("DB_NAME")
	config.DB.Password = os.Getenv("DB_PASSWORD")
	return config
}
