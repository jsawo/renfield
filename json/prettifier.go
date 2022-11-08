package json

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jsawo/renfield/cache"
	"github.com/jsawo/renfield/editor"
	"os"

	"github.com/jsawo/renfield/config"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type JSONFormatter struct {
	Ctx context.Context
}

func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}

// PrettifyJSON returns a prettified json string with a given indentation
func (j *JSONFormatter) PrettifyJSON(indent int, input string) string {
	if input == "" {
		return ""
	}

	cache.SaveCacheFile(input, "json_i")

	indentstring := fmt.Sprintf("%*s", indent, "")

	var result bytes.Buffer
	err := json.Indent(&result, []byte(input), "", indentstring)
	if err != nil {
		runtime.LogError(j.Ctx, err.Error())
		return err.Error()
	}

	cache.SaveCacheFile(result.String(), "json_o")

	return result.String()
}

func (j *JSONFormatter) GetLastCode() editor.EditorContent {
	contentIn, err := os.ReadFile(config.GetTempFilePath("json_i"))
	if err != nil {
		contentIn = []byte("-no preset content found-")
	}

	contentOut, _ := os.ReadFile(config.GetTempFilePath("json_o"))

	return editor.EditorContent{
		Input:  string(contentIn),
		Output: string(contentOut),
	}
}
