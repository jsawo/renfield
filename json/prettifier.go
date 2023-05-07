package json

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/jsawo/renfield/cache"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// PrettifyJSON returns a prettified json string with a given indentation
func (j *JSONTools) PrettifyJSON(indent int, input string) string {
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
