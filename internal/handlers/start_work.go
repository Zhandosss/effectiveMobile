package handlers

import (
	"context"
	"effectiveMobileTestProblem/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type StartWorkRequest struct {
	Id string `json:"id"`
}

// StartWork starts work
// @Summary Start work
// @Description Start work
// @ID startWork
// @Tags works
// @Accept json
// @Produce json
// @Param work body model.Work true "Work information"
// @Success 201 {object} map[string]string
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /work [post]
func (h *Handlers) StartWork(c echo.Context) error {
	var req model.Work
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msgf("Failed to bind request")
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request. Check request body"})
	}
	log.Info().Msgf("For request with id: %s. StartWork request body: %+v", c.Response().Header().Get(echo.HeaderXRequestID), req)
	req.StartTime = time.Now()

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	id, err := h.WorkService.StartWork(ctx, &req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to start work")
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to start work"})
	}
	log.Info().Msgf("For request with id: %s. Work started with id %s.", c.Response().Header().Get(echo.HeaderXRequestID), id)
	return c.JSON(http.StatusCreated, StartWorkRequest{Id: id})

}
