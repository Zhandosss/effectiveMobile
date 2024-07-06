package handlers

import (
	"context"
	"effectiveMobileTestProblem/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (h *Handlers) UpdateUserById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request"})
	}

	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request"})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()
	err := h.UserService.UpdateUserById(ctx, id, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Error updating user"})
	}

	return c.JSON(http.StatusOK, nil)
}

func (h *Handlers) UpdateUserByPassport(c echo.Context) error {
	passport := c.Param("passport")
	if passport == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request"})
	}

	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request"})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()
	err := h.UserService.UpdateUserByPassport(ctx, passport, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Error updating user"})
	}

	return c.JSON(http.StatusOK, nil)
}
