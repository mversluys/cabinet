package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	c := cabinetConfiguration()
	m, err := mameGetMachines(c.Mame, c.Romset...)
	if err != nil {
		log.Fatal("mame scan failed", err)

	}

	log.Println("mame", m)

	app := NewApp(c, m)
	http.HandleFunc("/video", app.ServeVideo)
	go func() {
		http.ListenAndServe(":8080", nil)
	}()

	err = wails.Run(&options.App{
		Title:  "cabinet",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		LogLevel:         logger.DEBUG,
		WindowStartState: options.Fullscreen,
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		//Debug: options.Debug{ OpenInspectorOnStartup: true, },
	})

	if err != nil {
		log.Println("Error:", err.Error())
	}
}
