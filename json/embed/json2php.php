<?php
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
