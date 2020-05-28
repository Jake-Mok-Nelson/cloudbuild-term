package main

import (
	"fmt"

	"github.com/Jake-Mok-Nelson/cloudbuild-term/internal/config"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")                   // name of config file (without extension)
	viper.SetConfigType("yaml")                     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.cloudbuilder-term") // call multiple times to add many search paths
	viper.AddConfigPath(".")                        // optionally look for config in the working directory
	var configuration config.Configuration
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config file, %s", err))
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	viper.SetDefault("Projects", nil)
	viper.SetDefault("Theme", map[string]string{"BackgroundColour": "black", "ForgroundColour": "white"})

	println(configuration.Projects[1].Name)
}
