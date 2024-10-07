package main

import (
	"context"
	"net/http"
	"path/filepath"
)

type App struct {
	config   *cabinetConfig
	machines map[string]machine
	ctx      context.Context
}

func NewApp(config *cabinetConfig, machines map[string]machine) *App {
	return &App{
		config:   config,
		machines: machines,
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetRomset() []string {
	return a.config.Romset
}

func (a *App) GetMachines() map[string]machine {
	return a.machines
}

func (a *App) ServeVideo(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		http.Error(w, "Name parameter is required", http.StatusBadRequest)
		return
	}
	path := filepath.Join(a.config.Mame, "VideoSnaps", name+".mp4")
	http.ServeFile(w, r, path)
}

func (a *App) Play(rom string) {
	mamePlay(a.ctx, a.config.Mame, rom)
	bringAppToForeground()
	//runtime.EventsEmit(a.ctx, "focus")
}
