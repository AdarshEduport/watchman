package db

import (
	"database/sql"
	"log"
	"github.com/go-sql-driver/mysql"
)

func MyNewSQLStorage(cf mysql.Config)(db *sql.DB,e error ){
	
   status,err := sql.Open("mysql",cf.FormatDSN())
   if err!=nil{
	log.Fatal("DB Error ",err)
   }
   return status,err

}