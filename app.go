package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

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
}

// PrettifyJSON return a prettified json string with a given indentation
func (a *App) PrettifyJSON(indent int, input string) string {
	if input == "" {
		return ""
	}

	indentstring := fmt.Sprintf("%*s", indent, "")

	var result bytes.Buffer
	err := json.Indent(&result, []byte(input), "", indentstring)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return err.Error()
	}

	return result.String()
}
