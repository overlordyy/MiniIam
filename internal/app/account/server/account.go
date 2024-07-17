package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveAccount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "SaveAccount endpoint"})
}
