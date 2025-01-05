package game

import (
	"fmt"
	"rltest/pkg/app"
	"rltest/pkg/food"
	rl "rltest/pkg/raylib"
	"rltest/pkg/snake"
	"strconv"
)

const (
	StatePlay = iota
	StatePause
	StateLose
	StateWin
)

var frameCounter int = 0

type Game struct {
	MaxSnakeLength int
	SquareSize     int
	PlayerScore    int
	State          int
	Offset         rl.Vector2
	App            *app.App
}

func InitGame(app *app.App) *Game {
	game := &Game{}

	frameCounter = 0

	game.DefaultValues()

	game.App = app

	return game
}

func (game *Game) DefaultValues() {
	game.MaxSnakeLength = 10
	game.SquareSize = 31
	game.State = StatePlay
	frameCounter = 0
}

func (game *Game) GetOffset() {
	game.Offset.X = float32(game.App.ScreenWidth % game.SquareSize)
	game.Offset.Y = float32(game.App.ScreenHeight % game.SquareSize)
}

func (game *Game) DrawGameGrid() {
	for i := 0; i < game.App.ScreenWidth/game.SquareSize+1; i++ {
		rl.DrawLineV(
			rl.Vector2{
				X: float32(game.SquareSize*i) + game.Offset.X/2,
				Y: game.Offset.Y / 2,
			},
			rl.Vector2{
				X: float32(game.SquareSize*i) + game.Offset.X/2,
				Y: float32(game.App.ScreenHeight) - game.Offset.Y/2,
			},
			rl.LIGHTGRAY,
		)
	}

	for i := 0; i < game.App.ScreenHeight/game.SquareSize+1; i++ {
		rl.DrawLineV(
			rl.Vector2{
				X: game.Offset.X / 2,
				Y: float32(game.SquareSize*i) + game.Offset.Y/2,
			},

			rl.Vector2{
				X: float32(game.App.ScreenWidth) - game.Offset.X/2,
				Y: float32(game.SquareSize*i) + game.Offset.Y/2,
			},
			rl.LIGHTGRAY,
		)
	}
}

func (game *Game) RestartGame(player *snake.Snake) {
	game.DefaultValues()
	game.GetOffset()

	for i := 0; i < game.MaxSnakeLength; i++ {
		player.Segments[i] = rl.Vector2{
			X: game.Offset.X / 2,
			Y: game.Offset.Y / 2,
		}
	}

	for i := 0; i < player.Length; i++ {
		player.SnakePos[i] = rl.Vector2Zero()
	}

	player.Length = snake.InitSnakeLength
	player.Speed = rl.Vector2{X: float32(game.SquareSize), Y: 0}
}

func collisionWithScreen(head rl.Vector2, scrW, scrH int, offset rl.Vector2) (game_over bool) {
	game_over = false

	if (head.X > (float32(scrW) - offset.X)) ||
		(head.Y > (float32(scrH) - offset.Y)) ||
		(head.X < 0 || head.Y < 0) {
		game_over = true
		fmt.Printf(
			"Snake hits the wall:\nHead:\t\t% 4.2f :% 4.2f\nWall max:\t% 4.2f :% 4.2f\n",
			head.X,
			head.Y,
			(float32(scrW) - offset.X), (float32(scrH) - offset.Y),
		)
	}
	return
}

func (game *Game) UpdateGame(player *snake.Snake, fruit *food.Food) {
	if game.State == StateLose || game.State == StateWin {
		if rl.IsKeyPressed(rune(rl.KEY_ENTER)) {
			game.RestartGame(player)
		}
		return
	}

	if rl.IsKeyPressed('P') {
		game.State = (game.State + 1) % 2
	}

	if game.State == StatePlay {
		player.RotateSnake()
		player.Movement(frameCounter)

		if collisionWithScreen(player.Segments[0], game.App.ScreenWidth, game.App.ScreenHeight, game.Offset) || player.AteItself() {
			game.State = StateLose
		}

		SpawnFood(fruit, player, game.App, game)

		if game.State != StateWin && SnakeCollidesFood(player, fruit) {
			fruit.Active = false
		}

		frameCounter++

		if player.Length == game.MaxSnakeLength {
			game.State = StateWin
		}
	}
}

func (game *Game) DrawGame(player *snake.Snake, fruit *food.Food) {
	const (
		msgOver  string = "PRESS [ENTER] TO PLAY AGAIN"
		msgPause string = "GAME PAUSED"
		msgScore string = "SCORE:"
	)

	rl.BeginDrawing()
	rl.ClearBackground(rl.RAYWHITE)

	switch game.State {
	case StateLose:
		rl.DrawText(
			msgScore,
			rl.MeasureText(msgScore, 20)/2,
			rl.GetScreenHeight()/2-30,
			20,
			rl.GRAY,
		)
		rl.DrawText(strconv.Itoa(game.PlayerScore),
			rl.MeasureText(msgScore, 20)*2,
			rl.GetScreenHeight()/2-30,
			20,
			rl.GRAY,
		)
		rl.DrawText(msgOver,
			rl.GetScreenWidth()/2-rl.MeasureText(msgOver, 20)/2,
			rl.GetScreenHeight()/2-50,
			20,
			rl.GRAY,
		)

	case StateWin:
		rl.DrawText(
			"YOU WIN!",
			rl.GetScreenWidth()/2-rl.MeasureText("YOU WIN!", 20)/2,
			rl.GetScreenHeight()/2-50,
			20,
			rl.GRAY,
		)

	case StatePause:
		rl.DrawText(
			msgPause,
			game.App.ScreenWidth/2-rl.MeasureText(msgPause, 40)/2,
			game.App.ScreenHeight/2-40, 40, rl.GRAY,
		)
	}

	game.DrawGameGrid()
	player.DrawSnake()
	fruit.DrawFood()

	rl.EndDrawing()
}

func (game *Game) UnloadGame() {
	// TODO: Unload resources cleaning game
}
