package postgres

import (
	"errors"

	"github.com/ibrahimakbar31/comment-api-go/core/model/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

//DB1Connect initial connection DB1
func DB1Connect() (*db.DB1, error) {
	var err error
	db1Host := viper.GetString(viper.GetString("Env") + ".DB1.Host")
	db1Port := viper.GetString(viper.GetString("Env") + ".DB1.Port")
	db1User := viper.GetString(viper.GetString("Env") + ".DB1.User")
	db1Password := viper.GetString(viper.GetString("Env") + ".DB1.Password")
	db1DBName := viper.GetString(viper.GetString("Env") + ".DB1.DBName")
	db1Conn, err := gorm.Open("postgres", "host="+db1Host+" port="+db1Port+" user="+db1User+" dbname="+db1DBName+" password="+db1Password+" sslmode=disable")
	if err != nil {
		return &db.DB1{db1Conn}, errors.New("DB_CONNECTION_ERROR")
	}
	db1Conn.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	return &db.DB1{db1Conn}, err
}
