package json

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/jsawo/renfield/cache"
	"github.com/jsawo/renfield/config"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/sys"
)

func (j *JSONTools) Filter(filter string) string {
	input, _ := cache.ReadCacheFile("json_i")

	currentProject := config.GetCurrentProject()
	currentProject.JSONTools.Filter = filter
	config.UpdateProject(currentProject)

	if filter == "" {
		return j.prettifyJSON(2, input)
	}

	gronResult := j.filterJSON(input, filter)

	cache.SaveCacheFile(gronResult, "json_o")

	return gronResult
}

func (j *JSONTools) filterJSON(input, filter string) string {
	if input == "" {
		return ""
	}

	if j.gronModule == nil {
		j.gronModule = j.compileModule(j.GronWasmBytes)
	}

	// gron
	gronResult, err := j.runGronWasi(input)
	if err != nil {
		return err.Error()
	}

	// grep
	gronLines := strings.Split(gronResult, "\n")
	var greppedLines []string
	for _, line := range gronLines {
		if strings.Contains(line, filter) {
			greppedLines = append(greppedLines, line)
		}
	}

	greppedResult := strings.Join(greppedLines, "\n")

	// ungron
	ungronResult, err := j.runGronWasi(greppedResult, "--ungron")
	if err != nil {
		return err.Error()
	}

	return ungronResult
}

func (j *JSONTools) runGronWasi(input string, args ...string) (string, error) {
	ctx := context.Background()

	args = append([]string{"--"}, args...)

	stdin := bytes.NewBufferString(input)
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	config := wazero.NewModuleConfig().
		WithStdin(stdin).
		WithStdout(&stdout).
		WithStderr(&stderr).
		WithArgs(args...).
		WithName("gron.wasm")

	if _, err := (*j.wazeroRuntime).InstantiateModule(ctx, j.gronModule, config); err != nil {
		if exitErr, ok := err.(*sys.ExitError); ok && exitErr.ExitCode() != 0 {
			return "", fmt.Errorf("exit code %d", exitErr.ExitCode())
		} else if !ok {
			return "", err
		}
	}

	return stdout.String(), nil
}
