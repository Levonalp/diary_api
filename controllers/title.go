package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"diary_api/models"
	"diary_api/db"
)

func GetAllTitles(c *gin.Context) {
	titles, err := models.GetAllTitles(db.GetDB())
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": titles})
}

func GetTitleByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	title, err := models.GetTitleByID(db.GetDB(), uint(id))
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": title})
}

func CreateTitle(c *gin.Context) {
	var title models.Title
	err := c.ShouldBindJSON(&title)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	err = title.Create(db.GetDB())
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Title created successfully!"})
}

func UpdateTitle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var title models.Title
	err := c.ShouldBindJSON(&title)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	title.ID = uint(id)
	err = title.Update(db.GetDB())
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Title updated successfully!"})
}

