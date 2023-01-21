package main

import (
	"diary_api/db"
	"diary_api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	defer db.GetDB().Close()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/diaries", controllers.GetAllDiaries)
		v1.GET("/diary/:id", controllers.GetDiaryByID)
		v1.POST("/diary", controllers.CreateDiary)
		v1.PUT("/diary/:id", controllers.UpdateDiary)

		v1.GET("/titles", controllers.GetAllTitles)
		v1.GET("/title/:id", controllers.GetTitleByID)
		v1.POST("/title", controllers.CreateTitle)
		v1.PUT("/title/:id", controllers.UpdateTitle)

		v1.GET("/photos", controllers.GetAllPhotos)
		v1.GET("/photo/:id", controllers.GetPhotoByID)
		v1.POST("/photo", controllers.CreatePhoto)
		v1.PUT("/photo/:id", controllers.UpdatePhoto)
	}

	r.Run()
}
