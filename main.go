package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ibrahimakbar31/comment-api-go/middleware"
	"github.com/ibrahimakbar31/comment-api-go/router"
	"github.com/ibrahimakbar31/comment-api-go/service/postgres"
	"github.com/spf13/viper"
)

func init() {
	setEnv()
	initConfigFile()
}

func main() {
	var err error
	var app middleware.App
	//var postgresDB postgres.DB
	//add db connect here
	app.DB1, err = postgres.ConnectDB1()
	if err != nil {
		fmt.Println("cannot connect DB: ", err)
		os.Exit(1)
	}
	err = app.DB1.DB.DB().Ping()
	if err != nil {
		println("ga konek")
		os.Exit(1)
	}

	isMigration := viper.GetBool("Migration")
	if isMigration == true {
		println("migrating table....")
		postgres.DB1Migration(app.DB1)
		println("migration done.")
	}
	router.InitAllRoutes(&app)
	port := viper.GetString("Port")
	fmt.Println("Server Running on Port:", port)
	http.ListenAndServe(":"+port, app.Router)
}

func setEnv() {
	getEnv := os.Getenv("GOCUSTOMENV")
	viper.Set("Env", "Development")
	if len(getEnv) > 0 {
		if getEnv == "production" {
			viper.Set("Env", "Production")
		} else if getEnv == "staging" {
			viper.Set("Env", "Staging")
		}
	}
	fmt.Println("Environment: " + viper.GetString("Env"))
}

func initConfigFile() {
	getFilename := os.Getenv("GOCONFIGFILENAME")
	if getFilename == "" {
		getFilename = "config"
	}
	viper.SetConfigName(getFilename)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("unable to read config: %v\n", err)
		os.Exit(1)
	}
}
