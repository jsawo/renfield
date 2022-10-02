package main

import (
	"context"

	"github.com/jsawo/renfield/beam"
	"github.com/jsawo/renfield/json"
	"github.com/jsawo/renfield/tinker"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	beam.StartServer(ctx)
}

func (a *App) OpenDirectoryDialog() string {
	result, _ := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	return result
}

func (a *App) PrettifyJSON(indent int, input string) string {
	return json.PrettifyJSON(a.ctx, indent, input)
}

func (a *App) GetLastJSON() string {
	return json.GetLastCode(a.ctx)
}

func (a *App) ExecuteTinkerCommand(input string) string {
	return tinker.ExecuteCommand(a.ctx, input)
}

func (a *App) SetProjectDir(projectDir string) {
	tinker.SetProjectDir(a.ctx, projectDir)
}

func (a *App) GetProjectDir() string {
	return tinker.GetProjectDir(a.ctx)
}

func (a *App) GetLastCode() string {
	return tinker.GetLastCode(a.ctx)
}
