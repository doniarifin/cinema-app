package routes

import (
	"cinema-app/docs"
	"cinema-app/internal/controller"
	"cinema-app/internal/middleware"
	model "cinema-app/internal/model"
	"cinema-app/internal/repository"
	"cinema-app/internal/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// c := gin.Context{}

	// apply middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	docs.SwaggerInfo.BasePath = ""

	repo := repository.NewUserRepository(db)
	//auth service
	srv := service.NewAuthService(repo)
	h := controller.NewAuthController(srv)
	//user service
	userSrv := service.NewUserService(repo)
	userController := controller.NewUserController(userSrv)

	//cinema
	cinemaRepo := repository.NewCinemaRepository(db)
	cinemaService := service.NewCinemaService(cinemaRepo)
	cinemaController := controller.NewCinemaController(cinemaService)

	//movie
	movieRepo := repository.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepo)
	movieController := controller.NewMovieControlller(movieService)

	//showtime
	showtimeRepo := repository.NewShowtimeRepository(db)
	showtimeService := service.NewShowtimeService(showtimeRepo)
	showtimeController := controller.NewShowtimeController(showtimeService)

	//seat
	seatRepo := repository.NewSeatRepository(db)
	seatSrv := service.NewSeatService(seatRepo)
	seatController := controller.NewSeatController(seatSrv)

	//transaction
	trxRepo := repository.NewTransactionRepository(db)
	trxSrv := service.NewTransactionService(trxRepo, db)
	trxController := controller.NewTransactionController(trxSrv)

	// Public Routes
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Protected Routes - Require JWT
	auth := r.Group("/api/v1")
	auth.Use(middleware.JWTMiddleware())
	auth.GET("/GetAllUsers", userController.GetAllUsers)
	auth.GET("/GetCurrentUser", userController.GetCurrentUser)
	auth.PUT("/Update/:id", userController.Update)
	auth.DELETE("/Delete/:id", middleware.Authorization(model.AdminRole.String()), userController.Delete)

	//cinema
	auth.GET("/cinemas", cinemaController.GetAll)
	auth.GET("/cinemas/:id", cinemaController.GetByID)
	auth.POST("/cinemas", cinemaController.Create)
	auth.PUT("/cinemas/:id", cinemaController.Update)
	auth.DELETE("/cinemas/:id", middleware.Authorization(model.AdminRole.String()), cinemaController.Delete)

	//movie
	auth.GET("/movies", movieController.GetAll)
	auth.GET("/movies/:id", movieController.GetByID)
	auth.POST("/movies", movieController.Create)
	auth.PUT("/movies/:id", movieController.Update)
	auth.DELETE("/movies/:id", middleware.Authorization(model.AdminRole.String()), movieController.Delete)

	//showtime
	auth.GET("/showtimes", showtimeController.GetAll)
	auth.GET("/showtimes/:id", showtimeController.GetByID)
	auth.POST("/showtimes", showtimeController.Create)
	auth.PUT("/showtimes/:id", showtimeController.Update)
	auth.DELETE("/showtimes/:id", middleware.Authorization(model.AdminRole.String()), showtimeController.Delete)

	//seat
	auth.GET("/seats", seatController.GetAll)
	auth.GET("/seats/:id", seatController.GetByID)
	auth.POST("/seats", seatController.Create)
	auth.PUT("/seats/:id", seatController.Update)
	auth.DELETE("/seats/:id", middleware.Authorization(model.AdminRole.String()), seatController.Delete)

	//trx
	auth.POST("/trx", trxController.CreateTransaction)
	auth.POST("/trx/:id/pay", trxController.MarkAsPaid)
}
