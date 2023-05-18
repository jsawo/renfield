package json

import (
	"bytes"
	"context"
	"fmt"
	"html"
	"regexp"
	"strings"

	"github.com/jsawo/renfield/cache"
	"github.com/tetratelabs/wazero"
)

// JSONToPHP converts a json string to a php array
func (j *JSONTools) JSONToPHP(input string) string {
	if input == "" {
		return ""
	}

	cache.SaveCacheFile(input, "json_i")

	result := j.convertToPHP(input)

	cache.SaveCacheFile(result, "json_o")

	return result
}

// PHPToJSON converts a php array to json string
func (j *JSONTools) PHPToJSON(input string) string {
	if input == "" {
		return ""
	}

	cache.SaveCacheFile(input, "json_i")

	result := j.convertToJSON(input)

	cache.SaveCacheFile(result, "json_o")

	return result
}

func (j *JSONTools) convertToPHP(jsonInput string) string {
	code := `<?php
$json = <<<'IDENTIFIER'
%s
IDENTIFIER;
		
	function printArray(array $arr, int $indent = 0, string $indentStr = "\t"): string {
		$outerPad = $indent;
		$innerPad = $indent + 1;
		$out = '[';
		if (count($arr) === 0) {
				$out .= ']';
		} else {
			$out .= PHP_EOL;
			foreach ($arr as $k => $v) {
				$padding = str_repeat($indentStr, $innerPad);
				$pattern = '%%s%%s => %%s,';
				if (is_array($v)) {
					$v = printArray($v, $innerPad, $indentStr);
				} else if (is_string($v)) {
					$pattern = '%%s%%s => "%%s",';
				} else if (is_null($v)) {
					$pattern = '%%s%%s => NULL,';
				}
		
				$out .= sprintf($pattern, $padding, is_int($k) ? $k : "\"$k\"", $v) . PHP_EOL;
			}
			$out .= str_repeat($indentStr, $outerPad) . ']';
		}

		return $out;
	}

	$arr = json_decode($json, true);
	var_export(printArray($arr, 0, '  '));
	`

	phpInput := fmt.Sprintf(code, jsonInput)

	result := j.runPHPWasi(phpInput)

	out := cleanUpPHPOutput(result)

	return out
}

func (j *JSONTools) convertToJSON(phpInput string) string {
	code := `<?php
	$arrayInput = %s;
		
	$result = json_encode($arrayInput, JSON_PRETTY_PRINT);
	var_export($result);
	`

	php := fmt.Sprintf(code, phpInput)

	result := j.runPHPWasi(php)

	out := cleanUpPHPOutput(result)

	return out
}

func (j *JSONTools) runPHPWasi(input string) string {
	ctx := context.Background()

	if j.phpModule == nil {
		j.phpModule = j.compileModule(j.PHPWasmBytes)
	}

	stdin := bytes.NewBufferString(input)
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	config := wazero.NewModuleConfig().
		WithStdin(stdin).
		WithStdout(&stdout).
		WithStderr(&stderr)

	_, err := (*j.wazeroRuntime).InstantiateModule(ctx, j.phpModule, config)
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
