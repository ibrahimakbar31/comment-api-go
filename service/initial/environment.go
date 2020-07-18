package initial

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//SetEnvironment to set environment
func SetEnvironment() {
	getEnv := os.Getenv("GOCUSTOMENV")
	viper.Set("Env", "Development")
	if len(getEnv) > 0 {
		if getEnv == "production" {
			viper.Set("Env", "Production")
			gin.SetMode("release")
		} else if getEnv == "staging" {
			viper.Set("Env", "Staging")
			gin.SetMode("release")
		}
	}
	fmt.Println("Environment: " + viper.GetString("Env"))
}
