package user

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "indexUser_OK",
	})
}

func Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "createUser_OK",
	})
}
