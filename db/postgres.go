package db

import (
	"effectiveMobileTestProblem/configs"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func NewPostgres(config *configs.DBConfig) *sqlx.DB {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Name)

	conn, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		log.Fatal().Msgf("Error connecting to database: %s", err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal().Msgf("Error pinging database: %s", err)
	}
	return conn
}
