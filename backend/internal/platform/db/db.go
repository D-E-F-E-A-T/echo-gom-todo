package db

import (
	"fmt"
	"log"
	"todo/cmd/app/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

func New() *gorm.DB {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration config.Configuration

	if vErr := viper.ReadInConfig(); vErr != nil {
		log.Fatalf("Error reading config file, %s", vErr)
	}
	vErr := viper.Unmarshal(&configuration)
	if vErr != nil {
		log.Fatalf("unable to decode into struct, %v", vErr)
	}
	dbConfig := "host=" + configuration.Database.dbUrl + " port=" + configuration.Database.dbPort + " user=" + configuration.Database.dbUser + " dbname=" + configuration.Database.dbName
	db, err := gorm.Open("postgres", dbConfig)
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	return db
}
