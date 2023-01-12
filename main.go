package main

import (
    "diary_api/controller"
    "diary_api/database"
    "diary_api/middleware"
    "diary_api/model"
    "fmt"
    "log"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
	
	
)

func main() {
    loadEnv()
    loadDatabase()
    serveApplication()
}

func loadEnv() {
    err := godotenv.Load(".env.local")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func loadDatabase() {
    database.Connect()
    database.Database.AutoMigrate(&model.User{})
    database.Database.AutoMigrate(&model.Entry{})
}

func serveApplication() {
    router := gin.Default()

    publicRoutes := router.Group("/auth")
    publicRoutes.POST("/register", controller.Register)
    publicRoutes.POST("/login", controller.Login)

    protectedRoutes := router.Group("/api")
    protectedRoutes.Use(middleware.JWTAuthMiddleware())
    protectedRoutes.POST("/entry", controller.AddEntry)
    protectedRoutes.GET("/entry", controller.GetAllEntries)
    protectedRoutes.PUT("/entry/:id", controller.UpdateEntry)
    protectedRoutes.DELETE("/entry/:id", controller.DeleteEntry)
    protectedRoutes.PUT("/user/:id", controller.UpdateUser)
    protectedRoutes.DELETE("/user/:id", controller.DeleteUser)

    router.Use(cors.Default())
    router.Run(":8000")
    fmt.Println("Server running on port 8000")
}
