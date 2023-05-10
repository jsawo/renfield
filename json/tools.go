package json

import (
	"context"

	"github.com/jsawo/renfield/cache"
	"github.com/jsawo/renfield/editor"
)

type JSONTools struct {
	Ctx           context.Context
	PHPWasmBytes  []byte
	GronWasmBytes []byte
	WasmCachePath string
}

func NewJSONTools(phpWasm, gronWasm []byte, wasmCachePath string) *JSONTools {
	return &JSONTools{
		PHPWasmBytes:  phpWasm,
		GronWasmBytes: gronWasm,
		WasmCachePath: wasmCachePath,
	}
}

func (j *JSONTools) GetLastCode() editor.EditorContent {
	contentIn, ok := cache.ReadCacheFile("json_i")
	if !ok {
		contentIn = "-no preset content found-"
	}

	contentOut, _ := cache.ReadCacheFile("json_o")

	return editor.EditorContent{
		Input:  contentIn,
		Output: contentOut,
	}
}
