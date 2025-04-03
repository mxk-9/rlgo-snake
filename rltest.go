package main

import (
	"embed"
	"rltest/internal/app"
	"rltest/internal/food"
	// "rltest/pkg/gui"
	"rltest/internal/player"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed assets/*
var content embed.FS

var frameCounter int = 0

func init() {
	rl.SetCallbackFunc(main)
}

func main() {
	a := app.AppInit()
	g := app.NewGame(a)

	g.Offset = a.Offset(g.SquareSize)
	g.FrameTick = 5

	snk := player.NewSnake(g.SquareSize, g.MaxSnakeLength, player.Green, g.Offset)
	food := food.NewFood(g.SquareSize, rl.Orange)

	g.RaiseWindow()

	// butt, err := gui.NewButton(
	// 	&content, "assets/some_button.png", rl.Vector2{X: 500, Y: 200}, 2,
	// 	func() {
	// 		snk.Color = rl.Magenta
	// 		snk.ColorBorder = rl.Red
	// 	},
	// )
	// if err != nil {
	// 	return
	// }

	for !rl.WindowShouldClose() {
		g.UpdateGame(snk, food)
		// butt.Press(rl.GetMousePosition(), gui.IsMouseLeftORTapPressed())
		// butt.Draw()
		g.DrawGame(snk, food)
	}

	rl.CloseWindow()
}
