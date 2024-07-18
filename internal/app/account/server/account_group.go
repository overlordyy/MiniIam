package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveAccountGroup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "SaveAccount endpoint"})
}

func DelAccountGroup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "SaveAccount endpoint"})
}

func UpdateAccountGroup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "SaveAccount endpoint"})
}

func FindAccountGroup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "SaveAccount endpoint"})
}
