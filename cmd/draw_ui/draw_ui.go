package main

import (
	"embed"
	"fmt"
	"rltest/pkg/gui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	//go:embed assets/*
	content embed.FS
)

func main() {
	rl.InitWindow(600, 600, "Creating UI")
	rl.SetWindowState(rl.FlagWindowResizable)
	rl.SetWindowMinSize(1, 1)

	myContainer, err := gui.NewConrainer(
		rl.Vector2{X: 59, Y: 50},
		rl.Vector2{X: 70, Y: 70},
		3, 3,
	)

	if err != nil {
		rl.SetTraceLogLevel(rl.LogError)
		rl.TraceLog(rl.LogError, fmt.Sprintf("%v", err))
		rl.CloseWindow()
	}

	err = myContainer.SetTexture(&content, "assets/dpad.png", 5, 0)
	if err != nil {
		rl.SetTraceLogLevel(rl.LogError)
		rl.TraceLog(rl.LogError, fmt.Sprintf("%v", err))
		rl.CloseWindow()
	}
	defer myContainer.Unload()

	camera := rl.NewCamera2D(rl.Vector2Zero(), rl.Vector2Zero(), 0.0, 1)

	upButton := gui.NewButton(rl.Vector2Zero(), rl.Vector2Zero())
	upButton.BindFuncOnPress(func() {
		rl.SetTraceLogLevel(rl.LogInfo)
		rl.TraceLog(rl.LogInfo, "Up")
	})

	upButton.BindDrawFuncOnPress(func() {
		if upButton.DrawIfPressed {
			myContainer.SetTextureFrame(2)
		} else {
			myContainer.SetTextureFrame(1)
		}

	})

	if err != nil {
		rl.SetTraceLogLevel(rl.LogError)
		rl.TraceLog(rl.LogError, fmt.Sprintf("%v", err))
		rl.CloseWindow()
	}

	myContainer.InsertItem(upButton, 0, 1)

	for !rl.WindowShouldClose() {
		upButton.Press(rl.GetMousePosition(), gui.IsMouseLeftORTapPressed())
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		myContainer.Draw()

		rl.BeginMode2D(camera)
		rl.EndMode2D()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
