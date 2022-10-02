package config

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

const (
	appDir      = "renfield"
	configFile  = "config.json"
	tempFileExt = ".tmp"
)

func Initialize() {
	dirPath := GetAppConfigDir()
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Printf("ERR: Error creating config directory: %v", err)
	}

	err = viper.WriteConfigAs(path.Join(dirPath, configFile))
	if err != nil {
		fmt.Printf("ERR: Failed write to config file: %v", err)
	}
}

func Save() {
	_ = viper.WriteConfig()
}

func Set(key string, value any) {
	viper.Set(key, value)

	Save()
}

func Get(key string) interface{} {
	return viper.Get(key)
}

func GetString(key string) string {
	val, _ := Get(key).(string)
	return val
}

func GetInt(key string) int {
	val, _ := Get(key).(int)
	return val
}

func GetBool(key string) bool {
	val, _ := Get(key).(bool)
	return val
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

	err := viper.ReadInConfig()
	if err != nil {
		var e *viper.ConfigFileNotFoundError
		if errors.As(err, &e) {
			fmt.Println("Config file missing, recreating")
		} else {
			fmt.Printf("ERR: Config file failed to load: %v", err)
		}

		Initialize()
	}
}
