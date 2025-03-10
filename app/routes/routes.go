package routes

import (
	"github.com/amdrx480/go-lms/app/middlewares"
	"github.com/amdrx480/go-lms/controllers/categories"
	"github.com/amdrx480/go-lms/controllers/chapters"
	"github.com/amdrx480/go-lms/controllers/courses"
	"github.com/amdrx480/go-lms/controllers/lessons"
	"github.com/amdrx480/go-lms/controllers/modules"
	"github.com/amdrx480/go-lms/controllers/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware   echo.MiddlewareFunc
	JWTMiddleware      echojwt.Config
	CategoryController categories.CategoryController
	ChapterController  chapters.ChapterController
	CourseController   courses.CourseController
	LessonController   lessons.LessonController
	ModuleController   modules.ModuleController
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

	// Chapter Routes
	ChapterRoutes := e.Group("/api/v1/chapters", echojwt.WithConfig(cl.JWTMiddleware))
	ChapterRoutes.Use(middlewares.VerifyToken)
	ChapterRoutes.GET("", cl.ChapterController.GetAll)                                 // Get all chapters
	ChapterRoutes.GET("/:id", cl.ChapterController.GetByID)                            // Get chapter by ID
	ChapterRoutes.POST("", cl.ChapterController.Create, middlewares.VerifyAdmin)       // Create new chapter
	ChapterRoutes.PUT("/:id", cl.ChapterController.Update, middlewares.VerifyAdmin)    // Update chapter by ID
	ChapterRoutes.DELETE("/:id", cl.ChapterController.Delete, middlewares.VerifyAdmin) // Soft delete chapter by ID

	// Course Routes
	courseRoutes := e.Group("/api/v1/courses", echojwt.WithConfig(cl.JWTMiddleware))
	courseRoutes.Use(middlewares.VerifyToken)
	courseRoutes.GET("", cl.CourseController.GetAllWithModule)                                  // Get all courses
	courseRoutes.GET("/:id", cl.CourseController.GetByID)                                       // Get course by ID
	courseRoutes.POST("", cl.CourseController.Create, middlewares.VerifyAdmin)                  // Create new course
	courseRoutes.PUT("/:id", cl.CourseController.Update, middlewares.VerifyAdmin)               // Update course by ID
	courseRoutes.DELETE("/:id", cl.CourseController.Delete, middlewares.VerifyAdmin)            // Soft delete course by ID
	courseRoutes.POST("/:id", cl.CourseController.Restore, middlewares.VerifyAdmin)             // Restore soft-deleted course
	courseRoutes.DELETE("/:id/force", cl.CourseController.ForceDelete, middlewares.VerifyAdmin) // Permanently delete course

	// Lesson Routes
	LessonRoutes := e.Group("/api/v1/lessons", echojwt.WithConfig(cl.JWTMiddleware))
	LessonRoutes.Use(middlewares.VerifyToken)
	LessonRoutes.GET("", cl.LessonController.GetAll)                                 // Get all chapters
	LessonRoutes.GET("/:id", cl.LessonController.GetByID)                            // Get chapter by ID
	LessonRoutes.POST("", cl.LessonController.Create, middlewares.VerifyAdmin)       // Create new chapter
	LessonRoutes.PUT("/:id", cl.LessonController.Update, middlewares.VerifyAdmin)    // Update chapter by ID
	LessonRoutes.DELETE("/:id", cl.LessonController.Delete, middlewares.VerifyAdmin) // Soft delete chapter by ID

	// Module Routes
	moduleRoutes := e.Group("/api/v1/modules", echojwt.WithConfig(cl.JWTMiddleware))
	moduleRoutes.Use(middlewares.VerifyToken)
	moduleRoutes.GET("", cl.ModuleController.GetAllWithChapter)                      // Get all modules
	moduleRoutes.GET("/:id", cl.ModuleController.GetByID)                            // Get module by ID
	moduleRoutes.POST("", cl.ModuleController.Create, middlewares.VerifyAdmin)       // Create new module
	moduleRoutes.PUT("/:id", cl.ModuleController.Update, middlewares.VerifyAdmin)    // Update module by ID
	moduleRoutes.DELETE("/:id", cl.ModuleController.Delete, middlewares.VerifyAdmin) // Soft delete module by ID
}
