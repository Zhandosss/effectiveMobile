package handlers

import (
	"context"
	"effectiveMobileTestProblem/internal/model"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CreateUserRequest struct {
	Id string `json:"id"`
}

// CreateUser creates a new user
// @Summary Create a new user
// @Tags users
// @Description Create a new user with name, surname, address, passport series and number
// @ID createUser
// @Accept json
// @Produce json
// @Param user body model.User true "User information"
// @Success 201 {object} CreateUserRequest
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user [post]
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

	if err := validatePassportData(passportData); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()
	id, err := h.UserService.CreateUser(ctx, &req)

	if errors.Is(err, model.ErrAlreadyExists) {
		log.Error().Msgf("User with passport series and number: %s already exists", req.PassportSeriesAndNumber)
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "User with this passport series and number already exists"})
	}

	if err != nil {
		log.Error().Err(err).Msg("Failed to create user")
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to create user"})
	}
	log.Info().Msgf("For request with id: %s. User created with id %s.", c.Response().Header().Get(echo.HeaderXRequestID), id)
	return c.JSON(http.StatusCreated, CreateUserRequest{Id: id})
}

func validatePassportData(passportData []string) error {
	if len(passportData) != 2 {
		log.Error().Msgf("Invalid passport series and number format. Passport data: %+v", passportData)
		return fmt.Errorf(`Invalid passport series and number format. Should be in next format: 1234 123456`)
	}
	if _, err := strconv.Atoi(passportData[0]); err != nil {
		log.Error().Msgf("Invalid passport series. Passport data: %+v", passportData[0])
		return fmt.Errorf("Invalid passport series. Passport series should be a number")
	}

	if _, err := strconv.Atoi(passportData[1]); err != nil {
		log.Error().Msgf("Invalid passport number. Passport data: %+v", passportData[1])
		return fmt.Errorf("Invalid passport number. Passport number should be a number")
	}
	return nil
}
