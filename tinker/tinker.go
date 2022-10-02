package tinker

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/jsawo/renfield/config"
)

// SetProjectDir saves a given project path
func SetProjectDir(ctx context.Context, projectDir string) {
	config.Set("laravel.project", projectDir)
	config.Save()
}

func GetProjectDir(ctx context.Context) string {
	return config.GetString("laravel.project")
}

func GetLastCode(ctx context.Context) string {
	content, err := os.ReadFile(config.GetTempFilePath("tinker"))
	if err != nil {
		return "-no preset content found-"
	}

	return string(content)
}

// ExecuteCommand executes a tinker command and returns the result
func ExecuteCommand(ctx context.Context, input string) string {
	if input == "" {
		return ""
	}

	projectDir := config.GetString("laravel.project")

	if projectDir == "" {
		return "Error - no project directory selected"
	}

	tempFile := config.GetTempFilePath("tinker")
	err := os.WriteFile(tempFile, []byte(input), 0660)
	if err != nil {
		fmt.Printf("ERR: Error writing to temp file: %v \n", err)
		os.Exit(1)
	}

	commandString := fmt.Sprintf("php artisan tinker < %q", tempFile)
	cmd := exec.Command("sh", "-c", commandString)
	cmd.Dir = projectDir

	out, _ := cmd.CombinedOutput()
	return string(out)
}
