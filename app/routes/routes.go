package routes

import (
	"github.com/amdrx480/go-lms/app/middlewares"
	"github.com/amdrx480/go-lms/controllers/categories"
	"github.com/amdrx480/go-lms/controllers/courses"
	"github.com/amdrx480/go-lms/controllers/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware   echo.MiddlewareFunc
	JWTMiddleware      echojwt.Config
	CategoryController categories.CategoryController
	CourseController   courses.CourseController
	UserController     users.UserController
}

func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	// User Routes
	userRoutes := e.Group("/api/v1/users")
	userRoutes.POST("/register", cl.UserController.Register)
	userRoutes.POST("/login", cl.UserController.Login)
	userRoutes.GET(
		"/profile",
		cl.UserController.GetUserProfile, // Get user profile
		echojwt.WithConfig(cl.JWTMiddleware),
		middlewares.VerifyToken,
	)

	// Category Routes
	categoryRoutes := e.Group("/api/v1/categories", echojwt.WithConfig(cl.JWTMiddleware))
	categoryRoutes.Use(middlewares.VerifyToken)
	categoryRoutes.GET("", cl.CategoryController.GetAll)                                 // Get all categories
	categoryRoutes.POST("", cl.CategoryController.Create, middlewares.VerifyAdmin)       // Create new category
	categoryRoutes.PUT("/:id", cl.CategoryController.Update, middlewares.VerifyAdmin)    // Update category by ID
	categoryRoutes.DELETE("/:id", cl.CategoryController.Delete, middlewares.VerifyAdmin) // Delete category by ID

	// Course Routes
	courseRoutes := e.Group("/api/v1/courses", echojwt.WithConfig(cl.JWTMiddleware))
	courseRoutes.Use(middlewares.VerifyToken)
	courseRoutes.GET("", cl.CourseController.GetAll)                                            // Get all courses
	courseRoutes.GET("/:id", cl.CourseController.GetByID)                                       // Get course by ID
	courseRoutes.POST("", cl.CourseController.Create, middlewares.VerifyAdmin)                  // Create new course
	courseRoutes.PUT("/:id", cl.CourseController.Update, middlewares.VerifyAdmin)               // Update course by ID
	courseRoutes.DELETE("/:id", cl.CourseController.Delete, middlewares.VerifyAdmin)            // Soft delete course by ID
	courseRoutes.POST("/:id", cl.CourseController.Restore, middlewares.VerifyAdmin)             // Restore soft-deleted course
	courseRoutes.DELETE("/:id/force", cl.CourseController.ForceDelete, middlewares.VerifyAdmin) // Permanently delete course
}
