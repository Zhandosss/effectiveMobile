package handlers

import (
	"context"
	"effectiveMobileTestProblem/internal/model"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"time"
)

type GetUsersRequest struct {
	Users          []*model.User  `json:"users"`
	PaginationInfo PaginationInfo `json:"pagination"`
}

// GetUserById returns user by id
// @Summary Get user by id
// @Tags users
// @Description Get user by id
// @ID getUserById
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/{id} [get]
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
	if errors.Is(err, model.ErrNotFound) {
		log.Error().Msgf("User with id %s not found", id)
		return c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found"})
	}

	if err != nil {
		log.Error().Err(err).Msg("Error getting user")
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Error getting user"})
	}

	log.Info().Msgf("For request with id: %s. User with id %s found", c.Response().Header().Get(echo.HeaderXRequestID), id)
	return c.JSON(http.StatusOK, user)
}

// GetUserByPassport returns user by passport
// @Summary Get user by passport
// @Tags users
// @Description Get user by passport
// @ID getUserByPassport
// @Accept json
// @Produce json
// @Param passport path string true "User passport series and number"
// @Success 200 {object} model.User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/passport/{passport} [get]
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
	if errors.Is(err, model.ErrNotFound) {
		log.Error().Msgf("User with passport series and number %s not found", passport)
		return c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found"})
	}
	if err != nil {
		log.Error().Err(err).Msg("Error getting user")
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Error getting user"})
	}

	log.Info().Msgf("For request with id: %s. User with passport series and number %s found", c.Response().Header().Get(echo.HeaderXRequestID), passport)
	return c.JSON(http.StatusOK, user)
}

// GetUsers returns users by filters
// @Summary Get users
// @Tags users
// @Description Get users by filters
// @ID getUsers
// @Accept json
// @Produce json
// @Param passport_series query string false "User passport series"
// @Param passport_number query string false "User passport number"
// @Param name query string false "User name"
// @Param surname query string false "User surname"
// @Param address query string false "User address"
// @Param per_page query string false "per page pagination parameter" default(10)
// @Param page query string false "page pagination parameter" default(1)
// @Success 200 {object} GetUsersRequest
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user [get]
func (h *Handlers) GetUsers(c echo.Context) error {
	var err error
	passportSeries := c.QueryParam("passport_series")
	passportNumber := c.QueryParam("passport_number")
	name := c.QueryParam("name")
	surname := c.QueryParam("surname")
	address := c.QueryParam("address")
	log.Info().Msgf("For request with id: %s. GetUsers request with next filters: passport_series: %s, passport_number: %s, name: %s, surname: %s, address: %s", c.Response().Header().Get(echo.HeaderXRequestID), passportSeries, passportNumber, name, surname, address)

	perPage := c.QueryParam("per_page")
	page := c.QueryParam("page")

	if perPage == "" {
		perPage = "10"
	}

	var perPageInt int

	if perPageInt, err = strconv.Atoi(perPage); err != nil {
		log.Error().Err(err).Msg("Invalid per_page parameter")
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid per_page parameter"})
	}

	if page == "" {
		page = "1"
	}

	var pageInt int

	if pageInt, err = strconv.Atoi(page); err != nil {
		log.Error().Err(err).Msg("Invalid page parameter")
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid page parameter"})

	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, "passport_series", passportSeries)
	ctx = context.WithValue(ctx, "passport_number", passportNumber)
	ctx = context.WithValue(ctx, "name", name)
	ctx = context.WithValue(ctx, "surname", surname)
	ctx = context.WithValue(ctx, "address", address)

	ctx = context.WithValue(ctx, "per_page", perPage)
	ctx = context.WithValue(ctx, "page", page)

	users, err := h.UserService.GetUsers(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error getting users")
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Error getting users"})
	}

	log.Info().Msgf("For request with id: %s. Users found", c.Response().Header().Get(echo.HeaderXRequestID))
	return c.JSON(http.StatusOK, GetUsersRequest{Users: users, PaginationInfo: PaginationInfo{Limit: perPageInt, Page: pageInt}})
}
