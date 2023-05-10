package json

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jsawo/renfield/cache"
	"github.com/jsawo/renfield/config"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"github.com/tetratelabs/wazero/sys"
)

func (j *JSONTools) Filter(filter string) string {
	fmt.Print("########## time start\n")
	start := time.Now()

	input, _ := cache.ReadCacheFile("json_i")

	currentProject := config.GetCurrentProject()
	currentProject.JSONTools.Filter = filter
	config.UpdateProject(currentProject)

	if filter == "" {
		return j.prettifyJSON(2, input)
	}

	stdin := bytes.NewBufferString(input)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	// gron
	gronResult, err := j.runGronWasi(stdin, &stdout, &stderr, false)
	if err != nil {
		return err.Error()
	}

	fmt.Printf("step gron took: %v\n", time.Since(start))

	// simple grep
	gronLines := strings.Split(gronResult, "\n")
	var greppedLines []string
	for _, line := range gronLines {
		if strings.Contains(line, filter) {
			greppedLines = append(greppedLines, line)
		}
	}

	if len(greppedLines) == 0 {
		return "{}"
	}

	greppedResult := strings.Join(greppedLines, "\n")

	fmt.Printf("step grep took: %v\n", time.Since(start))

	// ungron
	stdin = bytes.NewBufferString(greppedResult)
	stdout.Reset()
	stderr.Reset()
	ungronResult, err := j.runGronWasi(stdin, &stdout, &stderr, true)
	if err != nil {
		return err.Error()
	}

	fmt.Printf("step ungr took: %v\n", time.Since(start))

	cache.SaveCacheFile(ungronResult, "json_o")

	return ungronResult
}

func (j *JSONTools) runGronWasi(stdin, stdout, stderr *bytes.Buffer, ungron bool) (string, error) {
	ctx := context.Background()

	// configure compilation cache
	cache, err := wazero.NewCompilationCacheWithDir(j.WasmCachePath)
	if err != nil {
		return "", err
	}
	defer cache.Close(ctx)
	runtimeConfig := wazero.NewRuntimeConfig().WithCompilationCache(cache)

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntimeWithConfig(ctx, runtimeConfig)
	defer r.Close(ctx) // close everything this Runtime created

	config := wazero.NewModuleConfig().
		WithStdin(stdin).
		WithStdout(stdout).
		WithStderr(stderr).
		WithArgs("--").
		WithName("gron.wasm")

	if ungron {
		config = config.WithArgs("--ungron")
	}

	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	module, err := r.CompileModule(ctx, j.GronWasmBytes)
	if err != nil {
		return "", err
	}

	if _, err = r.InstantiateModule(ctx, module, config); err != nil {
		if exitErr, ok := err.(*sys.ExitError); ok && exitErr.ExitCode() != 0 {
			return "", fmt.Errorf("exit code %d", exitErr.ExitCode())
		} else if !ok {
			return "", err
		}
	}

	return stdout.String(), nil
}
