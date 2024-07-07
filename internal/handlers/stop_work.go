package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

// StopWork stops work by id
// @Summary Stop work by id
// @Description Stop work by id
// @ID stopWork
// @Tags works
// @Accept json
// @Produce json
// @Param id path string true "Work ID"
// @Success 200 {object} nil
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /work/{id} [delete]
func (h *Handlers) StopWork(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		log.Info().Msg("Cannot get work id from request")
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request. Provide work id"})
	}

	log.Info().Msgf("For request with id: %s. StopWork request with id: %s", c.Response().Header().Get(echo.HeaderXRequestID), id)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	ctx = context.WithValue(ctx, "end_time", time.Now())

	defer cancel()

	err := h.WorkService.StopWork(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to stop work")
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to stop work"})
	}

	log.Info().Msgf("For request with id: %s. Work with id %s stopped", c.Response().Header().Get(echo.HeaderXRequestID), id)
	return c.NoContent(http.StatusOK)
}
