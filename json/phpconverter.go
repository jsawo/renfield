package json

import (
	"bytes"
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/jsawo/renfield/cache"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"github.com/tetratelabs/wazero/sys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// JSONToPHP converts a json string to a php array
func (j *JSONTools) JSONToPHP(input string) string {
	if input == "" {
		return ""
	}

	cache.SaveCacheFile(input, "json_i")

	result, err := j.convertToPHP(input)
	if err != nil {
		runtime.LogError(j.Ctx, err.Error())
		return err.Error()
	}

	cache.SaveCacheFile(result, "json_o")

	return result
}

// PHPToJSON converts a php array to json string
func (j *JSONTools) PHPToJSON(input string) string {
	if input == "" {
		return ""
	}

	cache.SaveCacheFile(input, "json_i")

	result, err := j.convertToJSON(input)
	if err != nil {
		runtime.LogError(j.Ctx, err.Error())
		return err.Error()
	}

	cache.SaveCacheFile(result, "json_o")

	return result
}

func (j *JSONTools) runPHPWasi(stdin, stdout, stderr *bytes.Buffer) (string, error) {
	ctx := context.Background()

	// configure compilation cache
	// cache, err := wazero.NewCompilationCacheWithDir("./cache")
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
		WithName("php-cgi-8.2.0.wasm")

	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	module, err := r.CompileModule(ctx, j.PHPWasmBytes)
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

func (j *JSONTools) convertToPHP(jsonInput string) (string, error) {
	code := `<?php
$json = <<<'IDENTIFIER'
%s
IDENTIFIER;
		
	function printArray(array $arr, int $indent = 0, string $indentStr = "\t"): string {
			$outerPad = $indent;
			$innerPad = $indent + 1;
			$out = '[';
			if (count($arr) == 0) {
					$out .= ']';
			} else {
					$out .= PHP_EOL;
					foreach ($arr as $k => $v) {
							$padding = str_repeat($indentStr, $innerPad);
							$pattern = '%%s%%s => %%s,';
							if (is_array($v)) $v = printArray($v, $innerPad, $indentStr);
							else if (is_string($v)) $pattern = '%%s%%s => "%%s",';

							$out .= sprintf($pattern, $padding, is_int($k) ? $k : "\"$k\"", $v) . PHP_EOL;
					}
					$out .= str_repeat($indentStr, $outerPad) . ']';
			}

			return $out;
	}

	$arr = json_decode($json, true);
	var_export(printArray($arr, 0, '  '));
	`

	php := fmt.Sprintf(code, jsonInput)
	stdin := bytes.NewBufferString(php)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	result, err := j.runPHPWasi(stdin, &stdout, &stderr)
	if err != nil {
		return "", err
	}

	// remove headers
	content := strings.Split(result, "\r\n\r\n")
	if len(content) < 2 {
		return "", fmt.Errorf("could not find headers in output: %s", content[0])
	}

	// Remove surrounding single quotes
	re := regexp.MustCompile(`^'|'$`)
	out := re.ReplaceAllString(content[1], "")

	return out, nil
}

func (j *JSONTools) convertToJSON(phpInput string) (string, error) {
	code := `<?php
	// $arrayInput = eval('@s');
	$arrayInput = %s;
		
	$result = json_encode($arrayInput, JSON_PRETTY_PRINT);
	var_export($result);
	`

	php := fmt.Sprintf(code, phpInput)
	stdin := bytes.NewBufferString(php)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	result, err := j.runPHPWasi(stdin, &stdout, &stderr)
	if err != nil {
		return "", err
	}

	// remove headers
	content := strings.Split(result, "\r\n\r\n")
	if len(content) < 2 {
		return "", fmt.Errorf("could not find headers in output: %s", content[0])
	}

	// Remove surrounding single quotes
	re := regexp.MustCompile(`^'|'$`)
	out := re.ReplaceAllString(content[1], "")

	return out, nil
}
