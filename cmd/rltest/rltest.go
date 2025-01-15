package main

import (
	"rltest/internal/app"
	"rltest/internal/food"
	"rltest/internal/player"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var frameCounter int = 0

func main() {
	a := app.AppInit()
	g := app.NewGame(a)

	g.Offset = a.Offset(g.SquareSize)
	g.FrameTick = 5

	snk := player.NewSnake(g.SquareSize, g.MaxSnakeLength, player.Red, g.Offset)
	food := food.NewFood(g.SquareSize, rl.Orange)

	a.RaiseWindow()

	for !rl.WindowShouldClose() {
		g.UpdateGame(snk, food)
		g.DrawGame(snk, food)
	}

	rl.CloseWindow()
}
