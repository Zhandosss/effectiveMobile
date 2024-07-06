package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func (h *Handlers) GetWorkById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		log.Error().Msg("Cannot get work id from request")
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request. Provide work id"})
	}
	log.Info().Msgf("For request with id: %s. GetWorkById request with id: %s", c.Response().Header().Get(echo.HeaderXRequestID), id)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	work, err := h.WorkService.GetWorkById(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Error getting work")
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Error getting work"})
	}
	log.Info().Msgf("For request with id: %s. Work with id %s found", c.Response().Header().Get(echo.HeaderXRequestID), id)
	return c.JSON(http.StatusOK, work)
}

func (h *Handlers) GetWorks(c echo.Context) error {
	user := c.QueryParams().Get("user")
	if user == "" {
		log.Error().Msg("Cannot get user id from request")
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request. Provide user id"})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, "end_time", time.Now())

	works, err := h.WorkService.GetWorks(ctx, user)
	if err != nil {
		log.Error().Err(err).Msg("Error getting works")
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Error getting works"})
	}

	log.Info().Msgf("For request with id: %s. Works for user %s found", c.Response().Header().Get(echo.HeaderXRequestID), user)
	return c.JSON(http.StatusOK, works)
}
