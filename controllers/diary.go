package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"diary_api/models"
	"diary_api/db"
)

func GetAllDiaries(c *gin.Context) {
	diaries, err := models.GetAllDiaries(db.GetDB())
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": diaries})
}

func GetDiaryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	diary, err := models.GetDiaryByID(db.GetDB(), uint(id))
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": diary})
}

func CreateDiary(c *gin.Context) {
	var diary models.Diary
	err := c.ShouldBindJSON(&diary)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	err = diary.Create(db.GetDB())
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Diary created successfully!"})
}

func UpdateDiary(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var diary models.Diary
	err := c.ShouldBindJSON(&diary)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	diary.ID = uint(id)
	err = diary.Update(db.GetDB())
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Diary updated successfully!"})
}

