package main

import (
	"errors"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")                   // name of config file (without extension)
	viper.SetConfigType("yaml")                     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.cloudbuilder-term") // call multiple times to add many search paths
	viper.AddConfigPath(".")                        // optionally look for config in the working directory
	err := validateConfig("$HOME/.cloudbuilder-term")
	if err != nil {
		panic(err)
	}

}

func validateConfig(configPath string) (err error) {
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		return errors.New("Fatal error reading config.yaml $HOME/.cloudbuilder-term/ or the current directory")
	}
	return nil
}