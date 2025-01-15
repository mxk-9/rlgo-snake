package app

import (
	"rltest/internal/food"
	"rltest/internal/player"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	statePlay uint8 = iota
	statePause
	stateLose
	stateWin
)

var emptyLevel = &Level{}

type Game struct {
	*App
	FrameTick      int32
	frameCounter   int32
	State          uint8
	SquareSize     int32
	MaxSnakeLength int32
	Offset         rl.Vector2
}

func NewGame(a *App) (game *Game) {
	game = &Game{}
	game.defaultValues()

	game.App = a
	return
}

func (g *Game) defaultValues() {
	g.FrameTick = 10
	g.MaxSnakeLength = 10
	g.SquareSize = 31
	g.State = statePlay
	g.frameCounter = 0
}

func (g *Game) RestartGame(snake *player.Snake) {
	g.defaultValues()

	for i := range g.MaxSnakeLength {
		snake.Segments[i] = rl.Vector2{X: g.Offset.X / 2, Y: g.Offset.Y / 2}
	}

	for i := range snake.Length {
		snake.SegPos[i] = rl.Vector2Zero()
	}

	snake.Length = player.InitSnakeLength
	snake.Speed = rl.Vector2{X: float32(g.SquareSize), Y: 0}
}

func (g *Game) UpdateGame(snake *player.Snake, food *food.Food) {
	if (g.State == stateLose || g.State == stateWin) && rl.IsKeyPressed(rl.KeyEnter) {
		g.RestartGame(snake)
	}

	if rl.IsKeyPressed('P') {
		g.State = (g.State + 1) % 2
	}

	if g.State == statePlay {
		g.frameCounter = snake.Move(g.frameCounter, g.FrameTick)

		SpawnFood(food, snake, g)

		if g.State != stateWin && SnakeCollidesFood(snake, food) {
			food.Active = false
		}

		g.frameCounter++

		if snake.Length == g.MaxSnakeLength {
			g.State = stateWin
		}

		if CollisionWithScreen(snake.Segments[0], g.ScreenWidth, g.ScreenHeight, g.Offset) || CollideYourself(snake) {
			g.State = stateLose
		}
	}
}

const (
	msgOver  string = "PRESS [ENTER] TO PLAY AGAIN"
	msgPause string = "GAME PAUSED"
)

func (g *Game) DrawGame(snake *player.Snake, food *food.Food) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	if g.State == stateLose {
		rl.DrawText(
			"SCORE:",
			rl.MeasureText("SCORE:", 20)/2,
			int32(rl.GetScreenHeight()/2-30),
			20,
			rl.Gray,
		)

		rl.DrawText(
			strconv.Itoa(int(snake.Length-player.InitSnakeLength)),
			rl.MeasureText("SCORE:", 20)*2,
			int32(rl.GetScreenHeight()/2-30),
			20,
			rl.Gray,
		)

		rl.DrawText(
			msgOver,
			int32(rl.GetScreenWidth()/2-int(rl.MeasureText(msgOver, 20)/2)),
			int32(rl.GetScreenHeight()/2-50),
			20,
			rl.Gray,
		)
		rl.EndDrawing()
		return
	}

	if g.State == stateWin {
		rl.DrawText(
			"YOU WIN!",
			int32(rl.GetScreenWidth()/2-int(rl.MeasureText("YOU WIN!", 20)/2)),
			int32(rl.GetScreenHeight()/2-50),
			20,
			rl.Gray,
		)
		rl.EndDrawing()
		return
	}

	emptyLevel.DrawGrid(g.ScreenWidth, g.ScreenHeight, g.SquareSize, g.Offset)
	snake.Draw()
	food.DrawFood()

	if g.State == statePause {
		rl.DrawText(
			msgPause,
			g.ScreenWidth/2-rl.MeasureText(msgPause, 40)/2,
			g.ScreenHeight/2-40,
			40,
			rl.Gray,
		)
	}

	rl.EndDrawing()
}
