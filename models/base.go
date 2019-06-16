package models

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	// dbHost := os.Getenv("db_host")
	dbType := os.Getenv("db_type")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?parseTime=true", username, password, dbPort, dbName) //Build connection string
	fmt.Println(dbUri)

	conn, err := gorm.Open(dbType, dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{})

}

func GetDB() *gorm.DB {
	return db
}
