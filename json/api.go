package json

import (
	"context"

	"github.com/jsawo/renfield/cache"
	"github.com/jsawo/renfield/config"
	"github.com/jsawo/renfield/editor"
)

type JSONTools struct {
	Ctx context.Context
}

func NewJSONTools() *JSONTools {
	return &JSONTools{}
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

// Filter filters a json object using a given filter
func (j *JSONTools) Filter(filter string) string {
	input, _ := cache.ReadCacheFile("json_i")

	currentProject := config.GetCurrentProject()
	currentProject.JSONTools.Filter = filter
	config.UpdateProject(currentProject)

	if filter == "" {
		return prettifyJSON(2, input)
	}

	gronResult := filterJSON(input, filter)

	cache.SaveCacheFile(gronResult, "json_o")

	return gronResult
}

// JSONToPHP converts a json string to a php array
func (j *JSONTools) JSONToPHP(input string) string {
	if input == "" {
		return ""
	}

	cache.SaveCacheFile(input, "json_i")

	result := convertToPHP(input)

	cache.SaveCacheFile(result, "json_o")

	return result
}

// PHPToJSON converts a php array to json string
func (j *JSONTools) PHPToJSON(input string) string {
	if input == "" {
		return ""
	}

	cache.SaveCacheFile(input, "json_i")

	result := convertToJSON(input)

	cache.SaveCacheFile(result, "json_o")

	return result
}

// PrettifyJSON returns a prettified json string with a given indentation
func (j *JSONTools) PrettifyJSON(indent int, input string) string {
	if input == "" {
		return ""
	}

	cache.SaveCacheFile(input, "json_i")

	result := prettifyJSON(indent, input)

	cache.SaveCacheFile(result, "json_o")

	return result
}
