package json

import (
	"context"

	"github.com/jsawo/renfield/cache"
	"github.com/jsawo/renfield/editor"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

type JSONTools struct {
	Ctx           context.Context
	PHPWasmBytes  []byte
	GronWasmBytes []byte
	phpModule     wazero.CompiledModule
	gronModule    wazero.CompiledModule
	wazeroRuntime *wazero.Runtime
}

func NewJSONTools(phpWasm, gronWasm []byte, wasmCachePath string) *JSONTools {
	ctx := context.Background()

	cache, err := wazero.NewCompilationCacheWithDir(wasmCachePath)
	if err != nil {
		panic(err)
	}

	runtime := wazero.NewRuntimeWithConfig(ctx,
		wazero.NewRuntimeConfig().WithCompilationCache(cache),
	)

	wasi_snapshot_preview1.MustInstantiate(ctx, runtime)

	return &JSONTools{
		PHPWasmBytes:  phpWasm,
		GronWasmBytes: gronWasm,
		wazeroRuntime: &runtime,
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

func (j *JSONTools) compileModule(module []byte) wazero.CompiledModule {
	ctx := context.Background()

	compiledModule, err := (*j.wazeroRuntime).CompileModule(ctx, module)
	if err != nil {
		panic(err)
	}

	return compiledModule
}
