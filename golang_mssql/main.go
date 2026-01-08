package main

import (
	"fmt"
	"log"
	"src/golang_mssql/dbconfig"
	"src/golang_mssql/middleware"
	auth "src/golang_mssql/middleware/auth"
	products "src/golang_mssql/middleware/products"
	users "src/golang_mssql/middleware/users"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	dbconfig.Connection()
	err1 := godotenv.Load(".env")
	if err1 != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(static.Serve("/", static.LocalFile("templates", true)))
	router.Static("/assets", "./assets")

	router.POST("/auth/signin", auth.Login)
	router.POST("/auth/signup", auth.Register)

	authGuard := router.Group("/api")
	authGuard.Use(middleware.AuthMiddleware())
	{
		authGuard.GET("/getallusers", users.GetAllusers)
		authGuard.GET("/getuserid/:id", users.GetUserid)
		authGuard.PATCH("/changepassword/:id", users.ChangePassword)
		authGuard.PATCH("/updateprofile/:id", users.UpdateProfile)
		authGuard.PATCH("/uploadpicture/:id", users.UploadPicture)
		authGuard.PATCH("/mfa/activate/:id", auth.MfaActivate)
		authGuard.PATCH("/mfa/verifytotp/:id", auth.MfaVerifyotp)
	}

	router.GET("/products/list/:page", products.ProductList)
	router.GET("/products/search/:page/:key", products.ProductSearch)
	router.GET("/products/report", products.ProductPDFReport)
	router.GET("/chart", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		val, lab := products.GetSalesData()
		graph := products.CreateBarChart(val, lab)
		graph.Render(c.Writer)
	})

	host := "http://127.0.0.1"
	port := "5000"
	address := fmt.Sprintf("%s:%s", host, port)
	log.Print("Listening to ", address)

	if err := router.Run(":5000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
