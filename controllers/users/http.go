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

	accessToken, refreshToken, err := uc.userUseCase.Login(ctx, userInput.ToDomain())
	if err != nil || accessToken == "" || refreshToken == "" {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "invalid email or password", "")
	}

	tokenData := controllers.AuthTokenData{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "token created", tokenData)
}

func (uc *UserController) RefreshToken(c echo.Context) error {
	userRequest := request.RefreshRequest{}
	ctx := c.Request().Context()

	if err := c.Bind(&userRequest); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request body", "")
	}

	accessToken, err := uc.userUseCase.RefreshAccessToken(ctx, userRequest.RefreshToken)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "refresh token invalid or expired", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "token refreshed", accessToken)

}

func (uc *UserController) GetUserProfile(c echo.Context) error {
	user, err := uc.userUseCase.GetUserProfile(c.Request().Context())

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "user not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "user found", response.FromDomain(user))
}
