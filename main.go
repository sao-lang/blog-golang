package main

import (
	"blog/internal/config"
	"blog/internal/dto"
	"blog/internal/handlers"
	"blog/internal/middlewares"
	"blog/internal/models"
	"blog/internal/repositories"
	"blog/internal/services"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var globalConfig *config.Config

func setupConfig() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err.Error())
	}
	globalConfig = conf
}

func registerMiddlewares(router *gin.Engine) {
	router.Use(middlewares.Response())
	router.Use(middlewares.Logger())
}

func setupUserRouter(router *gin.Engine, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	authHandler := handlers.NewAuthHandler(userService)

	router.POST("/register", middlewares.Validate(&dto.CreateUserDTO{}), authHandler.Register)
	router.POST("/login", middlewares.Auth(), middlewares.Validate(&dto.CreateUserDTO{}), authHandler.Login)
}

func setupDB() *gorm.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", globalConfig.Database.User, globalConfig.Database.Password, globalConfig.Database.Host, globalConfig.Database.Port, globalConfig.Database.DatabaseName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	migrateDB(db)
	return db
}

func migrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})

	if err != nil {
		panic("Failed to migrate database")
	}
}

func setupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	registerMiddlewares(router)
	router.GET("/songs", handlers.GetSongs)

	setupUserRouter(router, db)
	return router
}

func main() {
	setupConfig()
	db := setupDB()

	router := setupRouter(db)

	router.Run("0.0.0.0:9089")
}
