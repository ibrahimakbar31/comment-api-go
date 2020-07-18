package initial

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

//GetConfigFile function to get config file
func GetConfigFile() error {
	getFilename := os.Getenv("GOCONFIGFILENAME")
	if getFilename == "" {
		getFilename = "config-example"
	}
	viper.SetConfigName(getFilename)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return errors.New("CANNOT_READ_CONFIG_FILE")
	}
	return err
}
