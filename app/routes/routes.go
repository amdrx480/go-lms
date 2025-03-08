package routes

import (
	"github.com/amdrx480/go-clean-architecture-hexagonal/app/middlewares"
	"github.com/amdrx480/go-clean-architecture-hexagonal/controllers/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc
	JWTMiddleware    echojwt.Config
	UserController   users.UserController
}

func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	userRoutes := e.Group("/api/v1/users")

	userRoutes.POST("/register", cl.UserController.Register)
	userRoutes.POST("/login", cl.UserController.Login)
	userRoutes.GET(
		"/profile",
		cl.UserController.GetUserProfile, // Get user profile
		echojwt.WithConfig(cl.JWTMiddleware),
		middlewares.VerifyToken,
	)

	something := e.Group("/api/v1/something", echojwt.WithConfig(cl.JWTMiddleware))
	something.Use(middlewares.VerifyToken)
}
