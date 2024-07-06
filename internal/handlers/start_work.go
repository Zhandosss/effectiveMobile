package handlers

import (
	"context"
	"effectiveMobileTestProblem/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (h *Handlers) StartWork(c echo.Context) error {
	var req model.Work
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request"})
	}
	req.StartTime = time.Now()

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	id, err := h.WorkService.StartWork(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]string{"id": id})

}
