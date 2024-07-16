package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var dbClient *sqlx.DB
var gormDB *gorm.DB

func InitPostgres() error {
	connStr := "user=postgres password=123456 dbname=postgres sslmode=disable"
	var err error
	dbClient, err = sqlx.Open("postgres", connStr)
	if err != nil {
		fmt.Println("init postgres err:", err)
		return err
	}
	defer dbClient.Close()
	err = dbClient.Ping()
	if err != nil {
		fmt.Println("init postgres errc:", err)
		return err
	}
	fmt.Println("init postgres ok")
	return nil
}
