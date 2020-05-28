package main

import (
	"errors"
	"testing"

	"github.com/spf13/viper"
)

func TestMissingBuildConfigGivesError(t *testing.T) {

	// INIT
	viper.SetConfigName("config")                   // name of config file (without extension)
	viper.SetConfigType("yaml")                     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.cloudbuilder-term") // call multiple times to add many search paths
	viper.AddConfigPath(".")

	got := validateConfig("$HOME/.cloudbuilder-term")
	want := errors.New("Fatal error reading config.yaml $HOME/.cloudbuilder-term/ or the current directory")

	t.Helper()
	if got.Error() != want.Error() {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
