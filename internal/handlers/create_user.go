package handlers

import (
	"context"
	"effectiveMobileTestProblem/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (h *Handlers) CreateUser(c echo.Context) error {
	var req model.User
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msgf("Failed to bind request")
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request. Check request body"})
	}
	log.Info().Msgf("For request with id: %s. CreateUser request body: %+v", c.Response().Header().Get(echo.HeaderXRequestID), req)
	if req.PassportSeriesAndNumber == "" {
		log.Error().Msg("Passport series and number is required")
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: `Passport series and number is required in next format: "1234 123456"`})
	}

	//TODO: Check if passport series  and unique
	passportData := strings.Split(req.PassportSeriesAndNumber, " ")
	log.Debug().Msgf("Passport data: %+v", passportData)

	if len(passportData) != 2 {
		log.Error().Msgf("Invalid passport series and number format. Passport data: %+v", passportData)
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: `Invalid passport series and number format. Should be in next format: "1234 123456"`})
	}

	if _, err := strconv.Atoi(passportData[0]); err != nil {
		log.Error().Msgf("Invalid passport series. Passport data: %+v", passportData[0])
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid passport series. Should be a number"})
	}

	if _, err := strconv.Atoi(passportData[1]); err != nil {
		log.Error().Msgf("Invalid passport number. Passport data: %+v", passportData[1])
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid passport number. Should be a number"})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()
	id, err := h.UserService.CreateUser(ctx, &req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create user")
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to create user"})
	}
	log.Info().Msgf("For request with id: %s. User created with id %s.", c.Response().Header().Get(echo.HeaderXRequestID), id)
	return c.JSON(http.StatusCreated, map[string]string{"id": id})
}
