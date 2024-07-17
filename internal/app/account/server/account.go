package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveAccount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "SaveAccount endpoint"})
}

func DelAccount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "SaveAccount endpoint"})
}

func UpdateAccount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "SaveAccount endpoint"})
}

func FindAccount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "SaveAccount endpoint"})
}
