package account

import (
	"fmt"
	"github.com/overlordyy/MiniIam/internal/app/account/config"
	"github.com/overlordyy/MiniIam/internal/app/account/db"
	"github.com/overlordyy/MiniIam/internal/app/account/redis"
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
}
