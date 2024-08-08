package controllers

import (
	"mini-cms-api/middlewares"
	"mini-cms-api/models"
	"mini-cms-api/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service services.UserService
}

func InitUserController(jwtOptions models.JWTOptions) UserController {
	return UserController{
		service: services.InitUserService(jwtOptions),
	}
}

func (uc *UserController) Register(c echo.Context) error {
	var registerReq models.RegisterRequest

	if err := c.Bind(&registerReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := registerReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	user, err := uc.service.Register(registerReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "registration failed",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.User]{
		Status:  "success",
		Message: "user registered",
		Data:    user,
	})
}

func (uc *UserController) Login(c echo.Context) error {
	var loginReq models.LoginRequest

	if err := c.Bind(&loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := loginReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	token, err := uc.service.Login(loginReq)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.Response[string]{
			Status:  "failed",
			Message: "invalid email or password",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "login success",
		Data:    token,
	})
}

func (uc *UserController) GetUserInfo(c echo.Context) error {
	claim, err := middlewares.GetUser(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "user not found",
		})
	}

	userID := strconv.Itoa(claim.ID)

	user, err := uc.service.GetUserInfo(userID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "user not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.User]{
		Status:  "success",
		Message: "user information",
		Data:    user,
	})
}
