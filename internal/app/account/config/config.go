package config

var (
	ServerConfig struct {
		Redis_ip     string
		Redis_port   string
		Redis_passwd string
		Redis_db     int
		db_ip        string
		db_port      string
		db_user      string
		db_d         string
	}
)

func InitConfig() error {
	ServerConfig.Redis_ip = "127.0.0.1"
	ServerConfig.Redis_port = "6379"
	ServerConfig.Redis_passwd = ""
	ServerConfig.Redis_db = 0
	return nil
}
