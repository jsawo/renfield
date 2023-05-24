package tinker

import (
	"context"
	"fmt"
	"os/exec"
	"syscall"
	"time"

	"github.com/jsawo/renfield/config"
)

func executeTinkerCommand(tempFile string, project config.ProjectConfig) (string, error) {
	timeout := time.Duration(config.AppConfig.TinkerTimeout) * time.Second
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
