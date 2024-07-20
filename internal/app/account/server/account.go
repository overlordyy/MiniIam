package server

import (
	"github.com/gin-gonic/gin"
	rsp_account "github.com/overlordyy/MiniIam/api/account"
	"github.com/overlordyy/MiniIam/internal/public/login"
	"net/http"
)

func SaveAccount(c *gin.Context) {
	resutlStr := login.AclStatusText(0)
	c.JSON(http.StatusOK, gin.H{"code_str": resutlStr, "message": "save account sueccss"})
}

func DelAccount(c *gin.Context) {
	resutlStr := login.AclStatusText(0)
	c.JSON(http.StatusOK, gin.H{"code_str": resutlStr, "message": "del account sueccss"})
}

func UpdateAccount(c *gin.Context) {
	resutlStr := login.AclStatusText(0)
	c.JSON(http.StatusOK, gin.H{"code_str": resutlStr, "message": "update account sueccss"})
}

func FindAccount(c *gin.Context) {
	rsp_Account := []rsp_account.Rsp_account{}
	Account := rsp_account.Rsp_account{}
	Account.Id = 1
	Account.GroupId = 1
	Account.Status = 1
	Account.UserId = "1"
	Account.UserName = "test"
	rsp_Account = append(rsp_Account, Account)
	resutlStr := login.AclStatusText(0)
	c.JSON(http.StatusOK, gin.H{"code_str": resutlStr, "message": "find account sueccss", "result": rsp_Account})
}
