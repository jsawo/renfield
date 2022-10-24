package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.design/x/clipboard"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	return &App{}
}

func (a *App) OpenDirectoryDialog() string {
	result, _ := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	return result
}

func (a *App) CopyToClipboard(data string) {
	clipboard.Write(clipboard.FmtText, []byte(data))
}
