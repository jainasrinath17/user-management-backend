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

func (uc *UserController) GetUsers(c echo.Context) error {
	users, err := uc.Service.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

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
