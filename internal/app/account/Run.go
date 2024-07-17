package account

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/overlordyy/MiniIam/internal/app/account/config"
	"github.com/overlordyy/MiniIam/internal/app/account/db"
	"github.com/overlordyy/MiniIam/internal/app/account/redis"
	"github.com/overlordyy/MiniIam/internal/app/account/server"
)

func Run() {
	err := config.InitConfig()
	if err != nil {

	}
	//build trpc

	//build redis
	redis.InitRedis()
	//build db
	//common.InitDB()
	err = db.InitPostgres()
	if err != nil {
		fmt.Println(err)
	}

	g := gin.Default()
	accountGrouop := g.Group("/api/miniiam/account")
	accountGrouop.POST("/saveAccount", server.SaveAccount)
	accountGrouop.POST("/findAccount", server.FindAccount)
	accountGrouop.POST("/delAccount", server.DelAccount)
	accountGrouop.POST("/updateAccount", server.UpdateAccount)

	g.Run(":8880")
}
