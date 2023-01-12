package controller

import (
    "diary_api/helper"
    "diary_api/model"
    "github.com/gin-gonic/gin"
    "net/http"
	"strconv"
)
	


func AddEntry(context *gin.Context) {
    var input model.Entry
    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := helper.CurrentUser(context)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    input.UserID = user.ID

    savedEntry, err := input.Save()

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GetAllEntries(context *gin.Context) {
    user, err := helper.CurrentUser(context)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"data": user.Entries})
}

func UpdateEntry(context *gin.Context) {
    entryId, err := strconv.Atoi(context.Params.ByName("id"))
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entry ID"})
        return
    }

    var input model.Entry
    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := helper.CurrentUser(context)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    entry, err := user.GetEntryById(entryId)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    entry.Title = input.Title
    entry.Content = input.Content

    updatedEntry, err := entry.Save()
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"data": updatedEntry})
}
func DeleteEntry(context *gin.Context) {
    entryId, err := strconv.Atoi(context.Params.ByName("id"))
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entry ID"})
        return
    }

    user, err := helper.CurrentUser(context)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    entry, err := user.GetEntryById(entryId)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := entry.Delete(); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "Entry deleted successfully"})
}
