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
	Id      string
	Name    string
	Path    string
	Tag     string
	Command string
	Tinker  TinkerConfig
}

type TinkerConfig struct {
	Tabs []editor.Tab
}

type Tag struct {
	Label string
	Color string
}

var appConfig AppConfig

const (
	appDir      = "renfield"
	configFile  = "config"
	tempFileExt = ".tmp"
)

func Initialize() {
	dirPath := GetAppConfigDir()
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Fprint(os.Stderr, "ERROR: Error creating config directory:", err.Error())
	}

	setDefaults()

	err = viper.SafeWriteConfig()
	if err != nil {
		fmt.Fprint(os.Stderr, "ERROR: Failed write to config file: \n", err.Error())
	}
}

func Save(config AppConfig) {
	appConfig = config
	viper.Set("currentproject", appConfig.Currentproject)
	viper.Set("projects", appConfig.Projects)
	viper.Set("tags", appConfig.Tags)

	err := viper.WriteConfig()
	if err != nil {
		fmt.Fprint(os.Stderr, "ERROR: Failed write to config file:", err.Error())
	}
}

func GetConfig() AppConfig {
	return appConfig
}

func GetFreshConfig() AppConfig {
	err := viper.Unmarshal(&appConfig)
	if err != nil {
		panic(fmt.Errorf("ERROR: unable to decode into struct: %w", err))
	}

	return appConfig
}

func GetAppConfigDir() string {
	userConfigDir, _ := os.UserConfigDir()
	return path.Join(userConfigDir, appDir)
}

func GetTempFilePath(tempName string) string {
	return path.Join(GetAppConfigDir(), tempName+tempFileExt)
}

func Load() {
	viper.SetConfigName(configFile)
	viper.SetConfigType("json")
	viper.AddConfigPath(GetAppConfigDir())

	appConfig = AppConfig{}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, initialize
			Initialize()
			return
		}

		fmt.Fprint(os.Stderr, "ERROR:", err.Error())
		os.Exit(2)
	}

	appConfig = GetFreshConfig()
}
