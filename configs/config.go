package configs

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
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
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", "8080")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Msgf("Error reading config file, %s", err)
	}

	config := &Config{}

	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatal().Msgf("Unable to decode into struct, %v", err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal().Msgf("Error loading .env file")
	}

	config.DB.Password = os.Getenv("DB_PASSWORD")
	return config
}
