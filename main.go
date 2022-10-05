package main

import (
	"context"
	"embed"

	"github.com/jsawo/renfield/beam"
	"github.com/jsawo/renfield/config"
	"github.com/jsawo/renfield/json"
	"github.com/jsawo/renfield/tinker"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

var (
	appService    *App
	tinkerService *tinker.Tinker
	jsonFormatter *json.JSONFormatter
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	config.Load()

	appService = NewApp()
	tinkerService = tinker.NewTinker()
	jsonFormatter = json.NewJSONFormatter()

	err := wails.Run(&options.App{
		Title:            "Renfield",
		Width:            1324,
		Height:           1024,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        appStartup,
		WindowStartState: options.Minimised,
		Bind: []interface{}{
			appService,
			tinkerService,
			jsonFormatter,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func appStartup(ctx context.Context) {
	propagateContext(ctx)
	beam.StartServer(ctx)
}

func propagateContext(ctx context.Context) {
	appService.ctx = ctx
	tinkerService.Ctx = ctx
	jsonFormatter.Ctx = ctx
}
