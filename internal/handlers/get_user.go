package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func (h *Handlers) GetUserById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		log.Error().Msg("Cannot get user id from request")
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request. Provide user id"})
	}
	log.Info().Msgf("For request with id: %s. GetUserById request with id: %s", c.Response().Header().Get(echo.HeaderXRequestID), id)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()
	user, err := h.UserService.GetUserById(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Error getting user")
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Error getting user"})
	}

	log.Info().Msgf("For request with id: %s. User with id %s found", c.Response().Header().Get(echo.HeaderXRequestID), id)
	return c.JSON(http.StatusOK, user)
}

func (h *Handlers) GetUserByPassport(c echo.Context) error {
	passport := c.Param("passport")
	if passport == "" {
		log.Error().Msg("Cannot get user passport from request")
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request. Provide user passport series and number"})
	}
	log.Info().Msgf("For request with id: %s. GetUserByPassport request with passport: %s", c.Response().Header().Get(echo.HeaderXRequestID), passport)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()
	user, err := h.UserService.GetUserByPassport(ctx, passport)
	if err != nil {
		log.Error().Err(err).Msg("Error getting user")
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Error getting user"})
	}

	log.Info().Msgf("For request with id: %s. User with passport series and number %s found", c.Response().Header().Get(echo.HeaderXRequestID), passport)
	return c.JSON(http.StatusOK, user)
}

func (h *Handlers) GetUsers(c echo.Context) error {
	passportSeries := c.QueryParam("passport_series")
	passportNumber := c.QueryParam("passport_number")
	name := c.QueryParam("name")
	surname := c.QueryParam("surname")
	address := c.QueryParam("address")
	log.Info().Msgf("For request with id: %s. GetUsers request with next filters: passport_series: %s, passport_number: %s, name: %s, surname: %s, address: %s", c.Response().Header().Get(echo.HeaderXRequestID), passportSeries, passportNumber, name, surname, address)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, "passport_series", passportSeries)
	ctx = context.WithValue(ctx, "passport_number", passportNumber)
	ctx = context.WithValue(ctx, "name", name)
	ctx = context.WithValue(ctx, "surname", surname)
	ctx = context.WithValue(ctx, "address", address)

	users, err := h.UserService.GetUsers(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error getting users")
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Error getting users"})
	}

	log.Info().Msgf("For request with id: %s. Users found", c.Response().Header().Get(echo.HeaderXRequestID))
	return c.JSON(http.StatusOK, users)
}
