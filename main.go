package main

import (
	"effectiveMobileTestProblem/configs"
	"effectiveMobileTestProblem/db"
	"effectiveMobileTestProblem/internal/handlers"
	"effectiveMobileTestProblem/internal/repository"
	"effectiveMobileTestProblem/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cfg := configs.Load()
	log.Info().Msg("Config loaded successfully")
	log.Debug().Msgf("Server configs: %+v", cfg.Server)
	//TODO: do not log the DB password
	log.Debug().Msgf("DB configs: %+v", cfg.DB)

	// Start the server

	conn := db.NewPostgres(cfg.DB)
	log.Info().Msg("DB connection established")
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal().Msgf("Error closing connection: %s", err)
		}
		log.Info().Msg("DB connection closed")
	}()

	userRep := repository.NewUser(conn)
	log.Debug().Msg("User repository created")
	workRep := repository.NewWork(conn)
	log.Debug().Msg("Work repository created")

	userServ := service.NewUser(userRep)
	log.Debug().Msg("User service created")
	workServ := service.NewWork(workRep)
	log.Debug().Msg("Work service created")

	e := echo.New()
	handlers.NewHandlers(e, userServ, workServ)
	log.Debug().Msg("Handlers created")

	log.Info().Msgf("Starting server on %s:%s", cfg.Server.Host, cfg.Server.Port)
	e.Logger.Fatal(e.Start(cfg.Server.Host + ":" + cfg.Server.Port))

}
