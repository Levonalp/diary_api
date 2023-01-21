package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"diary_api/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	diaryGroup := router.Group("/diary")
	{
		diaryGroup.GET("", controllers.GetAllDiaries)
		diaryGroup.GET("/:id", controllers.GetDiaryByID)
		diaryGroup.POST("", controllers.CreateDiary)
		diaryGroup.PUT("/:id", controllers.UpdateDiary)
		diaryGroup.DELETE("/:id", controllers.DeleteDiary)
	}

	titleGroup := router.Group("/title")
	{
		titleGroup.GET("", controllers.GetAllTitles)
		titleGroup.GET("/:id", controllers.GetTitleByID)
		titleGroup.GET("/diary/:diary_id", controllers.GetTitlesByDiaryID)
		titleGroup.POST("", controllers.CreateTitle)
		titleGroup.PUT("/:id", controllers.UpdateTitle)
		titleGroup.DELETE("/:id", controllers.DeleteTitle)
	}

	photoGroup := router.Group("/photo")
	{
		photoGroup.GET("", controllers.GetAllPhotos)
		photoGroup.GET("/:id", controllers.GetPhotoByID)
		photoGroup.GET("/diary/:diary_id", controllers.GetPhotosByDiaryID)
		photoGroup.POST("", controllers.CreatePhoto)
		photoGroup.PUT("/:id", controllers.UpdatePhoto)
		photoGroup.DELETE("/:id", controllers.DeletePhoto)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "404 page not found"})
	})

	return router
}
