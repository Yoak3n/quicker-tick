package main

import (
	"context"
	"embed"
	"quicker-tick/app"
	"quicker-tick/internal/database"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS
var ctx *context.Context

func main() {
	// Create an instance of the app structure
	database.InitDatabase()
	app := app.NewApp()
	go app.BoardcastCtx(ctx)
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "quicker-tick",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:        true,
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		Windows: &windows.Options{
			DisableWindowIcon: true,
		},
		OnStartup: app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
