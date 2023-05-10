package cache

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jsawo/renfield/config"
)

func SaveCacheFile(input, filename string) string {
	tempFile := config.GetTempFilePath(filename)
	path := filepath.Dir(tempFile)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0760)
		if err != nil {
			panic("can't create cache file: " + err.Error())
		}
	}

	err := os.WriteFile(tempFile, []byte(input), 0660)
	if err != nil {
		fmt.Printf("ERR: Error writing to temp file: %v \n", err)
		os.Exit(1)
	}

	return tempFile
}

func ReadCacheFile(filename string) (string, bool) {
	content, err := os.ReadFile(config.GetTempFilePath(filename))
	if err != nil {
		return "", false
	}

	return string(content), true
}
