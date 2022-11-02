package tinker

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/jsawo/renfield/config"
)

type Tinker struct {
	Ctx context.Context
}

func NewTinker() *Tinker {
	return &Tinker{}
}

func (t *Tinker) GetLastCode() string {
	content, err := os.ReadFile(config.GetTempFilePath("tinker"))
	if err != nil {
		return "User::factory()->make()"
	}

	return string(content)
}

// ExecuteCommand executes a tinker command and returns the result
func (t *Tinker) ExecuteCommand(input string) string {
	if input == "" {
		return ""
	}

	currentProject := config.GetConfig().Currentproject
	projectDir := config.GetConfig().Projects[currentProject].Path

	if projectDir == "" {
		return "Error - no project directory selected"
	}

	tempFile := config.GetTempFilePath("tinker")
	err := os.WriteFile(tempFile, []byte(input), 0660)
	if err != nil {
		fmt.Printf("ERR: Error writing to temp file: %v \n", err)
		os.Exit(1)
	}

	commandString := fmt.Sprintf("cat %q | sed -e 's/^<?php//' | php artisan tinker", tempFile)
	cmd := exec.Command("sh", "-c", commandString)
	cmd.Dir = projectDir

	out, _ := cmd.CombinedOutput()
	return string(out)
}
