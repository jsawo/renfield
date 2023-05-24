package json

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"html"
	"regexp"
	"strings"

	"github.com/jsawo/renfield/wasm"
	"github.com/tetratelabs/wazero"
)

//go:embed embed/json2php.php
var json2phpCode []byte

//go:embed embed/php2json.php
var php2jsonCode []byte

func convertToPHP(jsonInput string) string {
	code := string(json2phpCode)

	phpInput := fmt.Sprintf(code, jsonInput)

	result := runPHPWasi(phpInput)

	out := cleanUpPHPOutput(result)

	return out
}

func convertToJSON(phpInput string) string {
	code := string(php2jsonCode)

	php := fmt.Sprintf(code, phpInput)

	result := runPHPWasi(php)

	out := cleanUpPHPOutput(result)

	return out
}

func runPHPWasi(input string) string {
	stdin := bytes.NewBufferString(input)
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	config := wazero.NewModuleConfig().
		WithStdin(stdin).
		WithStdout(&stdout).
		WithStderr(&stderr)

	_, err := (*wasm.Runtime).InstantiateModule(context.Background(), wasm.PHPModule, config)
	if err != nil {
		return cleanUpPHPErrorOutput(stdout.String())
	}

	return stdout.String()
}

func cleanUpPHPOutput(output string) string {
	// remove headers
	content := strings.Split(output, "\r\n\r\n")
	if len(content) < 2 {
		return content[0]
	}

	// Remove surrounding single quotes
	re := regexp.MustCompile(`^'|'$`)
	out := re.ReplaceAllString(content[1], "")

	return out
}

func cleanUpPHPErrorOutput(output string) string {
	// remove html tags
	output = regexp.MustCompile("<[^>]*>").
		ReplaceAllString(output, "")

	// unescape strings
	return html.UnescapeString(output)
}
