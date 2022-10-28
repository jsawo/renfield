package config

import (
	"fmt"
	"os"
	"path"

	"github.com/jsawo/renfield/editor"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Currentproject string `mapstructure:"currentproject"`
	Projects       map[string]ProjectConfig
	Tags           []Tag
}

type ProjectConfig struct {
	Id     string
	Name   string
	Tag    string
	Type   string
	Tinker TinkerConfig
}

type TinkerConfig struct {
	Tabs []editor.Tab
}

type Tag struct {
	Label string
	Color string
}

const (
	appDir      = "renfield"
	configFile  = "config.json"
	tempFileExt = ".tmp"
)

func Initialize() {
	dirPath := GetAppConfigDir()
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Fprint(os.Stderr, "ERROR: Error creating config directory:", err.Error())
	}

	setDefaults()

	err = viper.WriteConfigAs(path.Join(dirPath, configFile))
	if err != nil {
		fmt.Fprint(os.Stderr, "ERROR: Failed write to config file:", err.Error())
	}
}

func Save() {
	_ = viper.WriteConfig()
}

func GetConfig() AppConfig {
	var conf AppConfig

	err := viper.Unmarshal(&conf)
	if err != nil {
		fmt.Fprint(os.Stderr, "ERROR: unable to decode into struct:", err.Error())
	}

	return conf
}

func GetAppConfigDir() string {
	userConfigDir, _ := os.UserConfigDir()
	return path.Join(userConfigDir, appDir)
}

func GetTempFilePath(tempName string) string {
	return path.Join(GetAppConfigDir(), tempName+tempFileExt)
}

func Load() {
	viper.SetConfigName("config")
	viper.AddConfigPath(GetAppConfigDir())

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, initialize
			Initialize()
			return
		}

		fmt.Fprint(os.Stderr, "ERROR:", err.Error())
		os.Exit(2)
	}
}
