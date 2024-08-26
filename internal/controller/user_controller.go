package controller

import (
	"net/http"
	"strconv"
	"user-management-backend/internal/model"
	"user-management-backend/internal/service"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Service service.UserService
}

func NewUserController(e *echo.Echo, userService service.UserService) {
	controller := &UserController{userService}

	e.GET("/users", controller.GetUsers)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)
	e.GET("/users/:id", controller.GetUserByID)
}

// GetUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Produce  json
// @Success 200 {array} model.User
// @Failure 500 {string} string "Internal Server Error"
// @Router /users [get]
func (uc *UserController) GetUsers(c echo.Context) error {
	users, err := uc.Service.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body model.User true "User object"
// @Success 201 {object} model.User
// @Failure 400 {string} string "Bad Request"
// @Router /users [post]
func (uc *UserController) CreateUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := uc.Service.CreateUser(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body model.User true "Updated user object"
// @Success 200 {object} model.User
// @Failure 400 {string} string "Bad Request"
// @Router /users/{id} [put]
func (uc *UserController) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	user := new(model.User)
	user.ID = uint(id)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := uc.Service.UpdateUser(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete user by ID
// @Tags users
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/{id} [delete]
func (uc *UserController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	if err := uc.Service.DeleteUser(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get user by user ID
// @Tags users
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/{id} [get]
func (uc *UserController) GetUserByID(c echo.Context) error {
	user := new(model.User)
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if user, err = uc.Service.GetUserByID(uint(num)); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}
