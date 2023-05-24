package wasm

import (
	"context"
	_ "embed"

	"github.com/jsawo/renfield/config"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed embed/php-cgi-8.2.0-slim.wasm
var phpWasmBytes []byte

//go:embed embed/gron.wasm
var gronWasmBytes []byte

var (
	Runtime    *wazero.Runtime
	PHPModule  wazero.CompiledModule
	GronModule wazero.CompiledModule
)

func InitializeService() {
	initializeRuntime()

	if GronModule == nil {
		GronModule = compileModule(gronWasmBytes)
	}

	if PHPModule == nil {
		PHPModule = compileModule(phpWasmBytes)
	}
}

func initializeRuntime() {
	ctx := context.Background()
	cache, err := wazero.NewCompilationCacheWithDir(config.GetWASMCachePath())
	if err != nil {
		panic(err)
	}

	runtime := wazero.NewRuntimeWithConfig(ctx,
		wazero.NewRuntimeConfig().
			WithCompilationCache(cache),
	)

	wasi_snapshot_preview1.MustInstantiate(ctx, runtime)

	Runtime = &runtime
}

func compileModule(module []byte) wazero.CompiledModule {
	ctx := context.Background()

	compiledModule, err := (*Runtime).CompileModule(ctx, module)
	if err != nil {
		panic(err)
	}

	return compiledModule
}
