package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Env_load() {
	// 環境変数GO_ENVは適当につけました
	if os.Getenv("GO_ENV") == "" {
		os.Setenv("GO_ENV", "development")
	}

	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

type Post struct {
	gorm.Model
	ImageURL     string
	CreatureName string
	Caption      string
	Address      string
	Latitude     float64
	Longitude    float64
	UserID       uint
	User         User
}

func main() {
	Env_load()
	dsn :=
		os.Getenv("DB_URL")

	//db接続
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// マイグレーション
	db.AutoMigrate(&User{}, &Post{})

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"post": "aa",
		})
	})

	router.Run(":8000") // 0.0.0.0:8000 でサーバーを立てます。
}
