package tinker

import (
	"context"
	"os"

	"github.com/jsawo/renfield/cache"
	"github.com/jsawo/renfield/config"
	"github.com/jsawo/renfield/editor"
)

type Tinker struct {
	Ctx            context.Context
	commandTimeout int
}

func NewTinker() *Tinker {
	return &Tinker{
		commandTimeout: config.AppConfig.TinkerTimeout,
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

	out, _ := executeTinkerCommand(tempFile, currentProject)

	cache.SaveCacheFile(out, "tinker/"+currentTab+"/out")

	return out
}
