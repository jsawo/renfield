package main

import (
	"context"

	"github.com/jsawo/renfield/config"
	"github.com/jsawo/renfield/project"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.design/x/clipboard"
)

type App struct {
	ctx    context.Context
	config config.Config
}

func NewApp() *App {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	return &App{
		config: config.AppConfig,
	}
}

func (a *App) OpenDirectoryDialog() string {
	result, _ := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	return result
}

func (a *App) CopyToClipboard(data string) {
	clipboard.Write(clipboard.FmtText, []byte(data))
}

func (a *App) GetConfig() config.Config {
	return a.config
}

func (a *App) SetCurrentProject(projectId string) {
	a.config.Currentproject = projectId
	a.saveConfig()
}

func (a *App) RemoveProject(projectId string) {
	projectCount := len(a.config.Projects)

	if projectCount == 1 {
		return
	}

	delete(a.config.Projects, projectId)

	for k := range a.config.Projects {
		a.config.Currentproject = k
		break
	}

	a.saveConfig()
}

func (a *App) CreateProject() string {
	id := project.GetNewId()
	newProject := config.ProjectConfig{
		Id:      id,
		Name:    "Unnamed",
		Tag:     "local",
		Command: config.DefaultCommand,
	}
	a.config.Projects[id] = newProject
	a.config.Currentproject = id

	a.saveConfig()
	return id
}

func (a *App) UpdateProjectSettings(projectId string, settings config.ProjectConfig) {
	a.config.Projects[projectId] = settings
	a.saveConfig()
}

func (a *App) saveConfig() {
	runtime.EventsEmit(a.ctx, "configUpdated", a.config)
	config.Save(a.config)
}
