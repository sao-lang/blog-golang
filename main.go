package main

import (
	"blog/internal/config"
	"blog/internal/dto"
	"blog/internal/handlers"
	"blog/internal/middlewares"
	"blog/internal/models"
	"blog/internal/repositories"
	"blog/internal/services"
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// "context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	// "net/http"
	// "os"
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

func setupPgSql() *gorm.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", globalConfig.PgSql.User, globalConfig.PgSql.Password, globalConfig.PgSql.Host, globalConfig.PgSql.Port, globalConfig.PgSql.DatabaseName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&models.User{})

	if err != nil {
		panic("Failed to migrate database")
	}
	return db
}

func setupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	registerMiddlewares(router)
	router.GET("/songs", handlers.GetSongs)

	setupUserRouter(router, db)
	return router
}

func setupMinIO() {

	// MinIO服务器的URL和端口
	endpoint := "localhost:9000"
	accessKeyID := "root"
	secretAccessKey := "12345678"
	useSSL := false

	// 创建一个MinIO客户端
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	// 创建一个上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 创建存储桶
	bucketName := "my-bucket"
	location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// 检查存储桶是否已经存在
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("Bucket %s already exists\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
	// 上传文件
	objectName := "1.txt"
	filePath := "./static/1.txt"
	contentType := "application/text"

	// 上传文件
	uploadInfo, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, uploadInfo.Size)

}

func main() {
	setupConfig()
	db := setupPgSql()
	setupMinIO()
	router := setupRouter(db)

	router.Run("0.0.0.0:10000")
}
