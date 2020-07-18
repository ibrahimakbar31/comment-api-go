package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ibrahimakbar31/comment-api-go/middleware"
	"github.com/ibrahimakbar31/comment-api-go/router"
	"github.com/ibrahimakbar31/comment-api-go/service/initial"
	"github.com/ibrahimakbar31/comment-api-go/service/postgres"
	"github.com/spf13/viper"
)

func init() {
	initial.SetEnvironment()
	err := initial.GetConfigFile()
	if err != nil {
		fmt.Printf("unable to read config: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	var err error
	var app middleware.App
	//add db connect here
	err = app.GetDB()
	if err != nil {
		fmt.Println("cannot connect DB: ", err)
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
