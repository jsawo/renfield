package json

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func prettifyJSON(indent int, input string) string {
	indentstring := fmt.Sprintf("%*s", indent, "")

	var result bytes.Buffer
	err := json.Indent(&result, []byte(input), "", indentstring)
	if err != nil {
		return err.Error()
	}

	return result.String()
}
