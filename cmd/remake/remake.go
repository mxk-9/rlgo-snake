package main

import (
	"rltest/pkg/app"
	"rltest/pkg/player"
	// rl "rltest/pkg/raylib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var frameCounter int = 0

func main() {
	a := app.AppInit()
	g := app.NewGame()
	level := &app.LevelMap{}
	// a.ScreenWidth = 1280
	// a.ScreenHeight = 800

	// rl.SetWindowSize(1280, 800)

	g.Offset = a.Offset(g.SquareSize)
	g.FrameTick = 5

	oo := player.NewSnake(g.SquareSize, g.MaxSnakeLength, player.Red, g.Offset)
	a.RaiseWindow()

	for !rl.WindowShouldClose() {

		oo.Move(&frameCounter, g.FrameTick)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		oo.Draw()
		level.DrawGrid(a.ScreenWidth, a.ScreenHeight, g.SquareSize, g.Offset)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
