package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	_driverFactory "github.com/amdrx480/go-lms/drivers"
	"github.com/amdrx480/go-lms/utils"

	_categoryUseCase "github.com/amdrx480/go-lms/businesses/categories"
	_categoryController "github.com/amdrx480/go-lms/controllers/categories"

	_chapterUseCase "github.com/amdrx480/go-lms/businesses/chapters"
	_chapterController "github.com/amdrx480/go-lms/controllers/chapters"

	_courseUseCase "github.com/amdrx480/go-lms/businesses/courses"
	_courseController "github.com/amdrx480/go-lms/controllers/courses"

	_documentUseCase "github.com/amdrx480/go-lms/businesses/documents"
	_documentController "github.com/amdrx480/go-lms/controllers/documents"

	_enrollmentUseCase "github.com/amdrx480/go-lms/businesses/enrollments"
	_enrollmentController "github.com/amdrx480/go-lms/controllers/enrollments"

	_lessonUseCase "github.com/amdrx480/go-lms/businesses/lessons"
	_lessonController "github.com/amdrx480/go-lms/controllers/lessons"

	_moduleUseCase "github.com/amdrx480/go-lms/businesses/modules"
	_moduleController "github.com/amdrx480/go-lms/controllers/modules"

	_userUseCase "github.com/amdrx480/go-lms/businesses/users"
	_userController "github.com/amdrx480/go-lms/controllers/users"

	_dbDriver "github.com/amdrx480/go-lms/drivers/mysql"

	_middleware "github.com/amdrx480/go-lms/app/middlewares"
	_routes "github.com/amdrx480/go-lms/app/routes"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type operation func(ctx context.Context) error

func main() {
	configDB := _dbDriver.DBConfig{
		DB_USERNAME: utils.GetConfig("DB_USERNAME"),
		DB_PASSWORD: utils.GetConfig("DB_PASSWORD"),
		DB_HOST:     utils.GetConfig("DB_HOST"),
		DB_PORT:     utils.GetConfig("DB_PORT"),
		DB_NAME:     utils.GetConfig("DB_NAME"),
	}

	db := configDB.InitDB()
	_dbDriver.MigrateDB(db)

	configJWT := _middleware.JWTConfig{
		SecretKey:       utils.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	configLogger := _middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	e := echo.New()

	// Middleware CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://192.168.58.1:3000"}, // *Ganti/Tambahkan domain sesuai kebutuhan
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	categoryRepo := _driverFactory.NewCategoryRepository(db)
	categoryUsecase := _categoryUseCase.NewCategoryUseCase(categoryRepo)
	categoryCtrl := _categoryController.NewCategoryController(categoryUsecase)

	chapterRepo := _driverFactory.NewChapterRepository(db)
	chapterUsecase := _chapterUseCase.NewChapterUseCase(chapterRepo)
	chapterCtrl := _chapterController.NewChapterController(chapterUsecase)

	courseRepo := _driverFactory.NewCourseRepository(db)
	courseUsecase := _courseUseCase.NewCourseUsecase(courseRepo)
	courseCtrl := _courseController.NewCourseController(courseUsecase)

	documentRepo := _driverFactory.NewDocumentRepository(db)
	documentUsecase := _documentUseCase.NewDocumentUseCase(documentRepo)
	documentCtrl := _documentController.NewDocumentController(documentUsecase)

	enrollmentRepo := _driverFactory.NewEnrollmentRepository(db)
	enrollmentUsecase := _enrollmentUseCase.NewEnrollmentUseCase(enrollmentRepo)
	enrollmentCtrl := _enrollmentController.NewEnrollmentController(enrollmentUsecase)

	lessonRepo := _driverFactory.NewLessonRepository(db)
	lessonUsecase := _lessonUseCase.NewLessonUseCase(lessonRepo)
	lessonCtrl := _lessonController.NewLessonController(lessonUsecase)

	moduleRepo := _driverFactory.NewModuleRepository(db)
	moduleUsecase := _moduleUseCase.NewModuleUseCase(moduleRepo)
	moduleCtrl := _moduleController.NewModuleController(moduleUsecase)

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUseCase.NewUserUseCase(userRepo, &configJWT)
	userCtrl := _userController.NewAuthController(userUsecase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware:   configLogger.Init(),
		JWTMiddleware:      configJWT.Init(),
		CategoryController: *categoryCtrl,
		ChapterController:  *chapterCtrl,
		CourseController:   *courseCtrl,
		DocumentController: *documentCtrl,
		EnrollmentCtrl:     *enrollmentCtrl,
		LessonController:   *lessonCtrl,
		ModuleController:   *moduleCtrl,
		UserController:     *userCtrl,
	}

	routesInit.RegisterRoutes(e)

	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"database": func(ctx context.Context) error {
			return _dbDriver.CloseDB(db)
		},
		"http-server": func(ctx context.Context) error {
			return e.Shutdown(context.Background())
		},
	})

	<-wait
}

// gracefulShutdown performs application shut down gracefully.
func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		// add any other syscalls that you want to be notified with
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Println("shutting down")

		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		// Do the operations asynchronously to save time
		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("cleaning up: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("%s: clean up failed: %s", innerKey, err.Error())
					return
				}

				log.Printf("%s was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}
