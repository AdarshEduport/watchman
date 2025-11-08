package main

import (
	"database/sql"
	"fmt"
	"log"
	"myapi/cmd/api"
	"myapi/config"
	"myapi/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.MyNewSQLStorage(mysql.Config{
		AllowNativePasswords: true,
		ParseTime:            true,
		DBName:               config.Envs.DBName,
		Passwd:               config.Envs.DBPassWd,
		Addr:                 config.Envs.DBAddr,
		Net:                  "tcp",
		User:                 config.Envs.DBUser,
		
	})
	if err!=nil{
		log.Fatal(err)
	}
	if db!=nil{
		initDb(db)
	}

	server := api.NewApiServer(fmt.Sprintf(":%s", config.Envs.Port), db)

	error := server.Run()
	if error != nil {
		log.Fatal("Error encountered :", error)

	}
}

func initDb(sql *sql.DB){
   err := sql.Ping()
   if err!=nil{
	log.Fatal("DB init failed ",err)
   }
   log.Println("DB initialized Successfully")
}
