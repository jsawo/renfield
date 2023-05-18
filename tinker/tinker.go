package tinker

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/jsawo/renfield/cache"
	"github.com/jsawo/renfield/config"
	"github.com/jsawo/renfield/editor"
)

type Tinker struct {
	Ctx            context.Context
	commandTimeout int
}

func NewTinker(timeout int) *Tinker {
	return &Tinker{
		commandTimeout: timeout,
	}
}

func (t *Tinker) GetLastCode(tabId string) editor.EditorContent {
	if tabId == "" {
		tabId = config.GetCurrentProject().Tinker.ActiveTab
	}

	contentIn, err := os.ReadFile(config.GetTempFilePath("tinker/" + tabId + "/in"))
	if err != nil {
		contentIn = []byte("User::factory()->make()")
	}

	contentOut, _ := os.ReadFile(config.GetTempFilePath("tinker/" + tabId + "/out"))

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
	currentTab := currentProject.Tinker.ActiveTab

	if currentProject.Path == "" {
		return "Error - no project directory selected"
	}

	tempFile := cache.SaveCacheFile(input, "tinker/"+currentTab+"/in")

	out, _ := t.executeTinkerCommand(tempFile, currentProject)

	cache.SaveCacheFile(out, "tinker/"+currentTab+"/out")

	return out
}

func (t *Tinker) executeTinkerCommand(tempFile string, project config.ProjectConfig) (string, error) {
	timeout := time.Duration(t.commandTimeout) * time.Second
	ctx, _ := context.WithTimeout(context.Background(), timeout)

	commandString := fmt.Sprintf("cat %q | %s", tempFile, project.Command)

	cmd := exec.CommandContext(ctx, "sh", "-c", commandString)
	cmd.Dir = project.Path
	cmd.WaitDelay = timeout
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true} // set pgid for child processes

	out, err := cmd.CombinedOutput()
	if ctx.Err() != nil {
		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	}

	if len(out) == 0 {
		return "", err
	}

	return string(out), nil
}
