package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)
type Config struct {
	DBUser string
	DBPassWd string
	DBAddr string
	DBName string
	Net string
	PublicHost string
	Port string

}
 var Envs = initConfig()

func initConfig()Config{
    godotenv.Load()
	return  Config{
		PublicHost:getEnv("PUBLIC_HOST","http://localhost") ,
		Port: getEnv("PORT","8080"),
		DBUser: getEnv("DB_USER","root"),
		DBPassWd: getEnv("DB_PASSWD","12345678"),
		DBName: getEnv("DB_NAME","watcher"),
		DBAddr: fmt.Sprintf("%s:%s",getEnv("DB_HOST","127.0.0.1"),getEnv("DB_PORT","3306")),
	}
}

func getEnv(key,fallback string) string{
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return fallback
}