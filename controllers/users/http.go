package users

import (
	"net/http"

	"github.com/amdrx480/go-lms/businesses/users"
	"github.com/amdrx480/go-lms/controllers"
	"github.com/amdrx480/go-lms/controllers/users/request"
	"github.com/amdrx480/go-lms/controllers/users/response"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase users.UseCase
}

func NewAuthController(authUC users.UseCase) *UserController {
	return &UserController{
		userUseCase: authUC,
	}
}

func (uc *UserController) Register(c echo.Context) error {
	userInput := request.UserRegister{}
	ctx := c.Request().Context()

	if err := c.Bind(&userInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	err := userInput.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	user, err := uc.userUseCase.Register(ctx, userInput.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "error when inserting data", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "user registered", response.FromDomain(user))
}

func (uc *UserController) Login(c echo.Context) error {
	userInput := request.UserLogin{}
	ctx := c.Request().Context()

	if err := c.Bind(&userInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	err := userInput.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	token, err := uc.userUseCase.Login(ctx, userInput.ToDomain())

	var isFailed bool = err != nil || token == ""

	if isFailed {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "invalid email or password", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "token created", token)
}

func (uc *UserController) GetUserProfile(c echo.Context) error {
	user, err := uc.userUseCase.GetUserProfile(c.Request().Context())

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "user not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "user found", response.FromDomain(user))
}
