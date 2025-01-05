package main

import (
	"fmt"
	"rltest/pkg/app"
	"rltest/pkg/food"
	"rltest/pkg/game"
	rl "rltest/pkg/raylib"
	"rltest/pkg/snake"
)

func main() {
	app := app.AppInit()
	fmt.Println("Init an app")
	game := game.InitGame(app)
	fmt.Println("Init a game")

	player := snake.NewSnake(2, game.SquareSize, game.MaxSnakeLength, game.Offset)
	fmt.Printf("Init a snake:% 4.0f:% 4.0f\n", player.Segments[0].X, player.Segments[0].Y)
	fruit := food.NewFood(game.SquareSize, rl.BLUE)
	fmt.Println("Init a fruit")

	for !rl.WindowShouldClose() {
		game.UpdateGame(player, fruit)
		game.DrawGame(player, fruit)
	}

	rl.CloseWindow()
}
