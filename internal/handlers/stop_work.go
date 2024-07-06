package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (h *Handlers) StopWork(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid id"})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	ctx = context.WithValue(ctx, "end_time", time.Now())

	defer cancel()

	err := h.WorkService.StopWork(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return c.NoContent(http.StatusOK)
}
