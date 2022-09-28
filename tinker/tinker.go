package tinker

import (
	"context"
	"fmt"
	"os/exec"
)

// ExecuteCommand executes a tinker command in a given folder and returns the result
func ExecuteCommand(ctx context.Context, projectDir, input string) string {
	if input == "" {
		return ""
	}

	if projectDir == "" {
		return "Error - no project directory selected"
	}

	commandString := fmt.Sprintf("echo %q | php artisan tinker", input)

	cmd := exec.Command("sh", "-c", commandString)

	cmd.Dir = projectDir

	out, _ := cmd.CombinedOutput()
	return string(out)
}
