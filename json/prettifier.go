package json

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jsawo/renfield/config"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// PrettifyJSON returns a prettified json string with a given indentation
func PrettifyJSON(ctx context.Context, indent int, input string) string {
	if input == "" {
		return ""
	}

	tempFile := config.GetTempFilePath("json")
	err := os.WriteFile(tempFile, []byte(input), 0660)
	if err != nil {
		fmt.Printf("ERR: Error writing to temp file: %v \n", err)
		os.Exit(1)
	}

	indentstring := fmt.Sprintf("%*s", indent, "")

	var result bytes.Buffer
	err = json.Indent(&result, []byte(input), "", indentstring)
	if err != nil {
		runtime.LogError(ctx, err.Error())
		return err.Error()
	}

	return result.String()
}

func GetLastCode(ctx context.Context) string {
	content, err := os.ReadFile(config.GetTempFilePath("json"))
	if err != nil {
		return "-no preset content found-"
	}

	return string(content)
}
