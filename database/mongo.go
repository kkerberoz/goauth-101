package database

import (
	"ginauth101/utils/config"
	"log"

	"github.com/goonode/mogo"
)

var mongoConnection *mogo.Connection = nil

//GetConnection is for get mongodb connection
func GetConnection() *mogo.Connection {
	if mongoConnection == nil {
		connectionString := config.EnvVar("DB_CONNECTION_STRING", "")
		dbName := config.EnvVar("DB_NAME", "")
		config := &mogo.Config{
			ConnectionString: connectionString,
			Database:         dbName,
		}
		mongoConnection, err := mogo.Connect(config)
		if err != nil {
			log.Fatal(err)
		} else {
			return mongoConnection
		}
	}
	return mongoConnection
}
