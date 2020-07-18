package router

import (
	"encoding/json"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ibrahimakbar31/comment-api-go/core/model"
	"github.com/ibrahimakbar31/comment-api-go/middleware"
	"github.com/ibrahimakbar31/comment-api-go/service/initial"
	"github.com/spf13/viper"
)

func GetAppTest(t *testing.T) *middleware.App {
	gin.SetMode(gin.TestMode)
	var app middleware.App
	initial.SetEnvironment()
	GetConfigFileTest(t)
	err := app.GetDB()
	if err != nil {
		t.Errorf(err.Error())
	}
	InitAllRoutes(&app)
	return &app
}

func GetConfigFileTest(t *testing.T) {
	viper.SetConfigName("config-example")
	viper.SetConfigType("json")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig()
	if err != nil {
		t.Errorf("CANNOT_READ_CONFIG_FILE")
	}
}

func MarshallingTestErrorResponse(t *testing.T, jsonData string, apiError *model.APIError) {
	byteData := []byte(jsonData)
	err := json.Unmarshal(byteData, apiError)
	if err != nil {
		t.Errorf(err.Error())
	}
}
