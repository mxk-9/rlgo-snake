package app

import (
	// rl "rltest/pkg/raylib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type App struct {
	MaxFps       int
	ScreenWidth  int
	ScreenHeight int
}

func AppInit() *App {
	app := &App{
		MaxFps:       60,
		ScreenWidth:  800,
		ScreenHeight: 800,
	}

	return app
}

func (app *App) RaiseWindow() {
	rl.InitWindow(int32(app.ScreenWidth), int32(app.ScreenWidth), "Bebra")
	rl.SetWindowMonitor(0)
	rl.SetTargetFPS(int32(app.MaxFps))
}

func (app *App) setAppWH() {}

func (app *App) Offset(square int) (offset rl.Vector2) {
	offset = rl.Vector2{
		X: float32(app.ScreenWidth % square),
		Y: float32(app.ScreenHeight % square),
	}
	return
}
