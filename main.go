package main

import (
	"net/http"
	"strconv"

	"mycrudapi/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	userService *service.UserService
}

func NewHandler(s *service.UserService) *Handler {
	return &Handler{userService: s}
}

// CreateUser handler
func (h *Handler) CreateUser(c echo.Context) error {
	u := new(service.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	createdUser := h.userService.CreateUser(u.Name, u.Email)
	return c.JSON(http.StatusCreated, createdUser)
}

// GetAllUsers handler
func (h *Handler) GetAllUsers(c echo.Context) error {
	users := h.userService.GetAllUsers()
	return c.JSON(http.StatusOK, users)
}

// GetUserByID handler
func (h *Handler) GetUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.userService.GetUserByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

// UpdateUser handler
func (h *Handler) UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	u := new(service.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	updatedUser, err := h.userService.UpdateUser(id, u.Name, u.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser handler
func (h *Handler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.userService.DeleteUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize UserService
	userService := service.NewUserService()

	// Initialize handler with userService dependency
	h := NewHandler(userService)

	// Routes
	e.POST("/users", h.CreateUser)     // Create a new user
	e.GET("/users", h.GetAllUsers)     // Get all users
	e.GET("/users/:id", h.GetUserByID) // Get a user by ID
	e.PUT("/users/:id", h.UpdateUser)  // Update a user by ID
	e.DELETE("/users/:id", h.DeleteUser) // Delete a user by ID

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
