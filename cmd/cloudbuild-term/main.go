package main

import (
	"fmt"
	"log"

	"github.com/Jake-Mok-Nelson/cloudbuild-term/internal/config"
	"github.com/Jake-Mok-Nelson/cloudbuild-term/internal/gui"
	"github.com/Jake-Mok-Nelson/cloudbuild-term/internal/projects"

	"github.com/jroimartin/gocui"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")                   // name of config file (without extension)
	viper.SetConfigType("yaml")                     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.cloudbuilder-term") // call multiple times to add many search paths
	viper.AddConfigPath(".")                        // optionally look for config in the working directory
	var configuration config.Configuration
	if err := viper.ReadInConfig(); err != nil {
		logrus.Warning("Coudldn't read config from $HOME/.cloudbuilder-term, assuming this is the first run and I'll create one")
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		panic(fmt.Errorf("Found some config but couldn't unmarshall it: , %v", err))
	}

	viper.SetDefault("Projects", nil)
	viper.SetDefault("Theme", map[string]string{"BackgroundColour": "black", "ForgroundColour": "white"})

	if viper.Get("Projects") == nil {
		panic("You don't have any projects configured")
	}

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true

	g.SetManagerFunc(gui.Layout)

	// Read the list of projects from config
	var projects []projects.Project
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("projects")
		if err != nil {
			return err
		}
		v.Clear()
		err = viper.UnmarshalKey("Projects", &projects)
		if err != nil {
			panic("Unable to unmarshal the projects from your config")
		}
		fmt.Fprintln(v, "ALL")
		for _, proj := range projects {
			fmt.Fprintln(v, proj.Name)
		}

		return nil
	})

	if err := gui.Keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
