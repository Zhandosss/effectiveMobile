package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func (h *Handlers) DeleteUserById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		log.Error().Msg("Cannot get user id from request")
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request. Provide user id"})
	}
	log.Info().Msgf("For request with id: %s. DeleteUserById request with id: %s", c.Response().Header().Get(echo.HeaderXRequestID), id)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	err := h.UserService.DeleteUserById(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting user")
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Error deleting user"})
	}

	log.Info().Msgf("For request with id: %s. User with id %s deleted", c.Response().Header().Get(echo.HeaderXRequestID), id)
	return c.JSON(http.StatusOK, nil)
}

func (h *Handlers) DeleteUserByPassport(c echo.Context) error {
	passport := c.Param("passport")
	if passport == "" {
		log.Error().Msg("Cannot get user passport from request")
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request"})
	}
	log.Info().Msgf("For request with id: %s. DeleteUserByPassport request with passport: %s", c.Response().Header().Get(echo.HeaderXRequestID), passport)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()
	err := h.UserService.DeleteUserByPassport(ctx, passport)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting user")
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Error deleting user"})
	}

	log.Info().Msgf("For request with id: %s. User with passport %s deleted", c.Response().Header().Get(echo.HeaderXRequestID), passport)
	return c.JSON(http.StatusOK, nil)
}
