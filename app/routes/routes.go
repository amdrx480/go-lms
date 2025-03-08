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
	AuthController   users.AuthController
}

func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	users := e.Group("/api/v1/users")

	users.POST("/register", cl.AuthController.Register)
	users.POST("/login", cl.AuthController.Login)

	something := e.Group("/api/v1/something", echojwt.WithConfig(cl.JWTMiddleware))
	something.Use(middlewares.VerifyToken)
}
