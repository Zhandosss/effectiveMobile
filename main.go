package main

import (
	"context"
	"effectiveMobileTestProblem/configs"
	"effectiveMobileTestProblem/db"
	"effectiveMobileTestProblem/internal/handlers"
	"effectiveMobileTestProblem/internal/repository"
	"effectiveMobileTestProblem/internal/service"
	"embed"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "effectiveMobileTestProblem/docs"
)

//go:embed db/migrations/*.sql
var embedMigrations embed.FS

//@title Time Tracker API
//@version 1.0
//@description This is a simple API for tracking time spent on work for users

// @host: localhost:8080
// @basePath /api/
func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cfg := configs.Load()
	log.Info().Msg("Config loaded successfully")
	log.Debug().Msgf("Server configs: %+v", cfg.Server)
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

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal().Msgf("Error setting dialect: %s", err)
	}

	if err := goose.Up(conn.DB, "db/migrations"); err != nil {
		log.Fatal().Msgf("Error applying migrations: %s", err)
	}

	userRep := repository.NewUser(conn)
	log.Debug().Msg("User repository created")
	workRep := repository.NewWork(conn)
	log.Debug().Msg("Work repository created")

	userServ := service.NewUser(userRep)
	log.Debug().Msg("User service created")
	workServ := service.NewWork(workRep)
	log.Debug().Msg("Work service created")

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	handlers.NewHandlers(e, userServ, workServ)
	log.Debug().Msg("Handlers created")

	log.Info().Msgf("Starting server on %s:%s", cfg.Server.Host, cfg.Server.Port)

	server := &http.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
		Handler:      e,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Msgf("Server shut down: %s", err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err)
	}
}
