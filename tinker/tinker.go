package tinker

import (
	"context"
	"fmt"
	"github.com/jsawo/renfield/cache"
	"github.com/jsawo/renfield/config"
	"github.com/jsawo/renfield/editor"
	"os"
	"os/exec"
)

type Tinker struct {
	Ctx context.Context
}

func NewTinker() *Tinker {
	return &Tinker{}
}

func (t *Tinker) GetLastCode() editor.EditorContent {
	contentIn, err := os.ReadFile(config.GetTempFilePath("tinker_i"))
	if err != nil {
		contentIn = []byte("User::factory()->make()")
	}

	contentOut, _ := os.ReadFile(config.GetTempFilePath("tinker_o"))

	return editor.EditorContent{
		Input:  string(contentIn),
		Output: string(contentOut),
	}
}

// ExecuteCommand executes a tinker command and returns the result
func (t *Tinker) ExecuteCommand(input string) string {
	if input == "" {
		return ""
	}

	currentProject := config.GetCurrentProject()

	if currentProject.Path == "" {
		return "Error - no project directory selected"
	}

	tempFile := cache.SaveCacheFile(input, "tinker_i")

	commandString := fmt.Sprintf("cat %q | sed -e 's/^<?php//' | %s", tempFile, currentProject.Command)

	cmd := exec.Command("sh", "-c", commandString)
	cmd.Dir = currentProject.Path

	out, _ := cmd.CombinedOutput()
	cache.SaveCacheFile(string(out), "tinker_o")

	return string(out)
}
