package routes

import (
	"github.com/amdrx480/angsana-boga/app/middlewares"
	"github.com/amdrx480/angsana-boga/controllers/categories"
	"github.com/amdrx480/angsana-boga/controllers/chapters"
	"github.com/amdrx480/angsana-boga/controllers/courses"
	"github.com/amdrx480/angsana-boga/controllers/documents"
	"github.com/amdrx480/angsana-boga/controllers/enrollments"
	"github.com/amdrx480/angsana-boga/controllers/lessons"
	"github.com/amdrx480/angsana-boga/controllers/modules"
	"github.com/amdrx480/angsana-boga/controllers/otp"
	"github.com/amdrx480/angsana-boga/controllers/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware   echo.MiddlewareFunc
	JWTMiddleware      echojwt.Config
	CategoryController categories.CategoryController
	ChapterController  chapters.ChapterController
	CourseController   courses.CourseController
	DocumentController documents.DocumentController
	EnrollmentCtrl     enrollments.EnrollmentController
	LessonController   lessons.LessonController
	ModuleController   modules.ModuleController
	UserController     users.UserController
	OTPController      otp.OTPController
}

func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	// User Routes
	userRoutes := e.Group("/api/v1")
	userRoutes.POST("/register", cl.UserController.Register)
	userRoutes.POST("/login", cl.UserController.Login)
	userRoutes.POST("/refresh-token", cl.UserController.RefreshToken)
	userRoutes.GET(
		"/profile",
		cl.UserController.GetUserProfile, // Get user profile
		echojwt.WithConfig(cl.JWTMiddleware),
		middlewares.VerifyToken,
	)

	// OTP Routes
	otpRoutes := e.Group("/api/v1/", echojwt.WithConfig(cl.JWTMiddleware))
	otpRoutes.Use(middlewares.VerifyToken)
	otpRoutes.POST("request-otp", cl.OTPController.RequestOTP)
	otpRoutes.POST("login-otp", cl.OTPController.LoginWithOTP)

	// Category Routes
	categoryRoutes := e.Group("/api/v1/categories", echojwt.WithConfig(cl.JWTMiddleware))
	categoryRoutes.Use(middlewares.VerifyToken)
	categoryRoutes.GET("", cl.CategoryController.GetAll)        // Get all categories
	categoryRoutes.POST("", cl.CategoryController.Create)       // Create new category
	categoryRoutes.PUT("/:id", cl.CategoryController.Update)    // Update category by ID
	categoryRoutes.DELETE("/:id", cl.CategoryController.Delete) // Delete category by ID

	// Chapter Routes
	ChapterRoutes := e.Group("/api/v1/chapters", echojwt.WithConfig(cl.JWTMiddleware))
	ChapterRoutes.Use(middlewares.VerifyToken)
	ChapterRoutes.GET("", cl.ChapterController.GetAll)        // Get all chapters
	ChapterRoutes.GET("/:id", cl.ChapterController.GetByID)   // Get chapter by ID
	ChapterRoutes.POST("", cl.ChapterController.Create)       // Create new chapter
	ChapterRoutes.PUT("/:id", cl.ChapterController.Update)    // Update chapter by ID
	ChapterRoutes.DELETE("/:id", cl.ChapterController.Delete) // Soft delete chapter by ID

	// Course Routes
	courseRoutes := e.Group("/api/v1/courses", echojwt.WithConfig(cl.JWTMiddleware))
	courseRoutes.Use(middlewares.VerifyToken)
	courseRoutes.GET("", cl.CourseController.GetAllWithModule)         // Get all courses
	courseRoutes.GET("/:id", cl.CourseController.GetByID)              // Get course by ID
	courseRoutes.POST("", cl.CourseController.Create)                  // Create new course
	courseRoutes.PUT("/:id", cl.CourseController.Update)               // Update course by ID
	courseRoutes.DELETE("/:id", cl.CourseController.Delete)            // Soft delete course by ID
	courseRoutes.POST("/:id", cl.CourseController.Restore)             // Restore soft-deleted course
	courseRoutes.DELETE("/:id/force", cl.CourseController.ForceDelete) // Permanently delete course

	// Document Routes
	DocumentRoutes := e.Group("/api/v1/documents", echojwt.WithConfig(cl.JWTMiddleware))
	DocumentRoutes.Use(middlewares.VerifyToken)
	DocumentRoutes.GET("", cl.DocumentController.GetAll)        // Get all documents
	DocumentRoutes.GET("/:id", cl.DocumentController.GetByID)   // Get document by ID
	DocumentRoutes.POST("", cl.DocumentController.Create)       // Create new document
	DocumentRoutes.PUT("/:id", cl.DocumentController.Update)    // Update document by ID
	DocumentRoutes.DELETE("/:id", cl.DocumentController.Delete) // Soft delete document by ID

	// Enrollment Routes
	EnrollmentRoutes := e.Group("/api/v1/enrollments", echojwt.WithConfig(cl.JWTMiddleware))
	EnrollmentRoutes.Use(middlewares.VerifyToken)
	EnrollmentRoutes.POST("", cl.EnrollmentCtrl.CreateEnrollmentCourse)                       //
	EnrollmentRoutes.GET("/:user_id/:course_id", cl.EnrollmentCtrl.GetEnrollmentByUserCourse) // Ambil enrollment berdasarkan user & course
	EnrollmentRoutes.GET("", cl.EnrollmentCtrl.GetAllEnrollmentCourseByUserID)                // Get lesson by ID

	// Lesson Routes
	LessonRoutes := e.Group("/api/v1/lessons", echojwt.WithConfig(cl.JWTMiddleware))
	LessonRoutes.Use(middlewares.VerifyToken)
	LessonRoutes.GET("", cl.LessonController.GetAll)        // Get all lessons
	LessonRoutes.GET("/:id", cl.LessonController.GetByID)   // Get lesson by ID
	LessonRoutes.POST("", cl.LessonController.Create)       // Create new lesson
	LessonRoutes.PUT("/:id", cl.LessonController.Update)    // Update lesson by ID
	LessonRoutes.DELETE("/:id", cl.LessonController.Delete) // Soft delete lesson by ID

	// Module Routes
	moduleRoutes := e.Group("/api/v1/modules", echojwt.WithConfig(cl.JWTMiddleware))
	moduleRoutes.Use(middlewares.VerifyToken)
	moduleRoutes.GET("", cl.ModuleController.GetAllWithChapter) // Get all modules
	moduleRoutes.GET("/:id", cl.ModuleController.GetByID)       // Get module by ID
	moduleRoutes.POST("", cl.ModuleController.Create)           // Create new module
	moduleRoutes.PUT("/:id", cl.ModuleController.Update)        // Update module by ID
	moduleRoutes.DELETE("/:id", cl.ModuleController.Delete)     // Soft delete module by ID
}

// Tambahkan ini untuk hanya role tertentu saja pada routes
// DocumentRoutes.POST("", cl.DocumentController.Create, middlewares.VerifyAdmin)       // Create new document
