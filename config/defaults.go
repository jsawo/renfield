package config

import (
	"github.com/jsawo/renfield/project"
	"github.com/spf13/viper"
)

func setDefaults() {
	defaultProjectId := project.GetNewId()
	viper.Set("projects."+defaultProjectId+".id", defaultProjectId)
	viper.Set("projects."+defaultProjectId+".type", "local")
	viper.Set("projects."+defaultProjectId+".name", "Unnamed")
	viper.Set("projects."+defaultProjectId+".path", "")
	viper.Set("projects."+defaultProjectId+".tag", "local")
	viper.Set("projects."+defaultProjectId+".tinker.tabs", []string{})
}
