package json

import (
	"bytes"
	"context"
	"strings"

	"github.com/jsawo/renfield/wasm"
	"github.com/tetratelabs/wazero"
)

func filterJSON(input, filter string) string {
	if input == "" {
		return ""
	}

	// gron
	gronResult, err := runGronWasi(input)
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

	if len(greppedLines) == 0 {
		return "{}"
	}

	greppedResult := strings.Join(greppedLines, "\n")

	// ungron
	ungronResult, err := runGronWasi(greppedResult, "--ungron")
	if err != nil {
		return err.Error()
	}

	return ungronResult
}

func runGronWasi(input string, args ...string) (string, error) {
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

	_, err := (*wasm.Runtime).InstantiateModule(context.Background(), wasm.GronModule, config)
	if err != nil {
		return "", err
	}

	return stdout.String(), nil
}
