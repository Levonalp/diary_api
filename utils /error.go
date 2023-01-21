package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func SendErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, Error{
		status,
		message,
	})
}

func HandleError(c *gin.Context, err error) {
	if err != nil {
		SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}
