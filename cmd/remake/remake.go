package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"rltest/pkg/app"
	"rltest/pkg/player"
)

var frameCounter int = 0

func main() {
	a := app.AppInit()
	g := app.NewGame()
	level := &app.LevelMap{}

	g.Offset = a.Offset(g.SquareSize)
	g.FrameTick = 5

	snk := player.NewSnake(g.SquareSize, g.MaxSnakeLength, player.Red, g.Offset)

	a.RaiseWindow()

	for !rl.WindowShouldClose() {

		snk.Move(&frameCounter, g.FrameTick)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		snk.Draw()
		level.DrawGrid(a.ScreenWidth, a.ScreenHeight, g.SquareSize, g.Offset)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
