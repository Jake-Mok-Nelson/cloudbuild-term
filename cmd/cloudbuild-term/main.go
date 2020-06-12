package main

import (
	"fmt"

	"github.com/Jake-Mok-Nelson/cloudbuild-term/internal/config"
	"github.com/Jake-Mok-Nelson/cloudbuild-term/internal/gui"
	"github.com/asaskevich/EventBus"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")                   // name of config file (without extension)
	viper.SetConfigType("yaml")                     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.cloudbuilder-term") // call multiple times to add many search paths
	viper.AddConfigPath(".")                        // optionally look for config in the working directory
	var configuration config.Configuration

	err := viper.Unmarshal(&configuration)
	if err != nil {
		panic(fmt.Errorf("Found some config but couldn't unmarshall it: , %v", err))
	}

	viper.SetDefault("Projects", nil)
	viper.SetDefault("Theme", map[string]string{"BackgroundColour": "black", "ForgroundColour": "white"})

	bus := EventBus.New()

	// Start the GUI
	gui.InitGUI(bus)

}
