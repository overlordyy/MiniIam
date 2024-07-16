package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/overlordyy/MiniIam/internal/app/account/config"
)

var (
	client redis.UniversalClient
)

func InitRedis() {
	fmt.Println(config.ServerConfig)
	client := redis.NewClient(&redis.Options{
		Addr:     config.ServerConfig.Redis_ip + ":" + config.ServerConfig.Redis_port,
		Password: config.ServerConfig.Redis_passwd,
		DB:       config.ServerConfig.Redis_db,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)
}
