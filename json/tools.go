package json

import (
	"context"
	"os"

	"github.com/jsawo/renfield/editor"

	"github.com/jsawo/renfield/config"
)

type JSONTools struct {
	Ctx           context.Context
	PHPWasmBytes  []byte
	WasmCachePath string
}

func NewJSONTools(phpWasm []byte, wasmCachePath string) *JSONTools {
	return &JSONTools{
		PHPWasmBytes:  phpWasm,
		WasmCachePath: wasmCachePath,
	}
}

func (j *JSONTools) GetLastCode() editor.EditorContent {
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
