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
	accountGrouop.Any("/saveAccount", server.SaveAccount)

	g.Run(":8880")
}
