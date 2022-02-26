package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"intrajasa-merchant-api-gateway/config"
	_ "github.com/go-sql-driver/mysql"
)

var connectionString string

func ConnectToDB(database string) (*gorm.DB, error) {
	db_connection := config.GoDotEnvVariable("DB_CONNECTION")
	db_host := config.GoDotEnvVariable("DB_HOST")
	db_port := config.GoDotEnvVariable("DB_PORT")
	db_database := config.GoDotEnvVariable("DB_DATABASE")
	db_username := config.GoDotEnvVariable("DB_USERNAME")
	db_password := config.GoDotEnvVariable("DB_PASSWORD")

	switch database {
	case "main":
		connectionString = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			db_username, db_password, db_host, db_port, db_database,
		)
		fmt.Println("run main access database")
	default:
		connectionString = fmt.Sprintf(
			"%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
			db_username, db_password, db_database,
		)
		fmt.Println("run default access database")
	}
	
	return gorm.Open(db_connection, connectionString)
}