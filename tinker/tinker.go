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

	appConfig := config.GetConfig()
	currentProjectId := appConfig.Currentproject
	currentProject := appConfig.Projects[currentProjectId]

	if currentProject.Path == "" {
		return "Error - no project directory selected"
	}

	tempFile := config.GetTempFilePath("tinker")
	err := os.WriteFile(tempFile, []byte(input), 0660)
	if err != nil {
		fmt.Printf("ERR: Error writing to temp file: %v \n", err)
		os.Exit(1)
	}

	commandString := fmt.Sprintf("cat %q | sed -e 's/^<?php//' | %s", tempFile, currentProject.Command)

	cmd := exec.Command("sh", "-c", commandString)
	cmd.Dir = currentProject.Path

	out, _ := cmd.CombinedOutput()
	return string(out)
}
