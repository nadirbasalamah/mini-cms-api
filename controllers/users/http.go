package users

import (
	"mini-cms-api/app/middlewares"
	"mini-cms-api/businesses/users"
	"mini-cms-api/controllers"
	"mini-cms-api/controllers/users/request"
	"mini-cms-api/controllers/users/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase users.UseCase
}

func NewUserController(userUC users.UseCase) *UserController {
	return &UserController{
		userUseCase: userUC,
	}
}

func (uc *UserController) Register(c echo.Context) error {
	registerReq := request.User{}

	if err := c.Bind(&registerReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	err := registerReq.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	user, err := uc.userUseCase.Register(registerReq.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "registration failed", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "user registered", response.FromDomain(user))
}

func (uc *UserController) Login(c echo.Context) error {
	loginReq := request.User{}

	if err := c.Bind(&loginReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	err := loginReq.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	token, err := uc.userUseCase.Login(loginReq.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "invalid email or password", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "login success", token)
}

func (uc *UserController) GetUserInfo(c echo.Context) error {
	claim, err := middlewares.GetUser(c)

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "user not found", "")
	}

	userID := strconv.Itoa(claim.ID)

	user, err := uc.userUseCase.GetUserInfo(userID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "user not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "user found", response.FromDomain(user))
}
