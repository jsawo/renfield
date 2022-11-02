package config

import (
	"github.com/jsawo/renfield/project"
	"github.com/spf13/viper"
)

func setDefaults() {
	defaultProjectId := project.GetNewId()
	project := "projects." + defaultProjectId

	viper.Set("currentproject", defaultProjectId)
	viper.Set(project+".id", defaultProjectId)
	viper.Set(project+".type", "local")
	viper.Set(project+".name", "Unnamed")
	viper.Set(project+".path", "")
	viper.Set(project+".tag", "local")
	viper.Set(project+".tinker.tabs", []string{})

	viper.Set("tags", []Tag{
		{Label: "local", Color: "green"},
		{Label: "testing", Color: "gray"},
		{Label: "development", Color: "blue"},
		{Label: "staging", Color: "orange"},
		{Label: "production", Color: "red"},
	})
}
