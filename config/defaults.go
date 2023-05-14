package config

import (
	"github.com/jsawo/renfield/project"
	"github.com/spf13/viper"
)

const DefaultCommand = "php artisan tinker"

func setDefaults() {
	defaultProjectId := project.GetNewId()
	projectDef := "projects." + defaultProjectId

	viper.Set("currentproject", defaultProjectId)
	viper.Set("tinkerTimeout", 2)
	viper.Set(projectDef+".id", defaultProjectId)
	viper.Set(projectDef+".command", DefaultCommand)
	viper.Set(projectDef+".name", "Unnamed")
	viper.Set(projectDef+".path", "")
	viper.Set(projectDef+".tag", "local")
	viper.Set(projectDef+".tinker.tabs", []string{})

	viper.Set("tags", []Tag{
		{Label: "local", Color: "green"},
		{Label: "testing", Color: "gray"},
		{Label: "development", Color: "blue"},
		{Label: "staging", Color: "orange"},
		{Label: "production", Color: "red"},
	})
}
