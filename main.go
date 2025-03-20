package main

import (
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

const (
	title   = "Dev Tools"
	version = "0.1"
	width   = 1080
	height  = 720
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:       fmt.Sprintf("%s - v%s", title, version),
		Width:       width,
		Height:      height,
		AlwaysOnTop: true,
		Debug:       options.Debug{OpenInspectorOnStartup: true},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 240, G: 247, B: 244, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
