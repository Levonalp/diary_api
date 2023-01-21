package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"diary_api/models"
	"diary_api/db"
)

func GetAllPhotos(c *gin.Context) {
	photos, err := models.GetAllPhotos(db.GetDB())
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": photos})
}

func GetPhotoByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	photo, err := models.GetPhotoByID(db.GetDB(), uint(id))
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": photo})
}

func CreatePhoto(c *gin.Context) {
	var photo models.Photo
	err := c.ShouldBindJSON(&photo)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	err = photo.Create(db.GetDB())
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Photo created successfully!"})
}

func UpdatePhoto(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var photo models.Photo
	err := c.ShouldBindJSON(&photo)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	photo.ID = uint(id)
	err = photo.Update(db.GetDB())
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully!"})
}
