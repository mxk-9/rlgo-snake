package game

import (
	"math/rand"
	"rltest/pkg/app"
	"rltest/pkg/food"
	rl "rltest/pkg/raylib"
	"rltest/pkg/snake"
)

func SnakeCollidesFood(snake *snake.Snake, food *food.Food) (collide bool) {
	snakeX := snake.Segments[0].X
	snakeY := snake.Segments[0].Y
	sSize := float32(snake.Size)
	foodX := food.Position.X
	foodY := food.Position.Y
	fSize := float32(food.Size)

	if (snakeX < (foodX + fSize)) && ((snakeX + sSize) > foodX) &&
		(snakeY < (foodX + fSize)) && ((snakeY + sSize) > foodY) {
		snake.Segments[snake.Length] = snake.SnakePos[snake.Length-1]
		snake.Length++
		collide = true
	}
	return
}

func SpawnFood(food *food.Food, snake *snake.Snake, app *app.App, game *Game) {
	if !food.Active {
		food.Active = true

		equ := true
		for i := 0; i < snake.Length; i++ {
			for equ = true; equ; equ = (food.Position.X == snake.Segments[i].X) || (food.Position.Y == snake.Segments[i].Y) {
				maxX := (app.ScreenWidth/game.SquareSize-1)*game.SquareSize + int(game.Offset.X/2)
				maxY := (app.ScreenHeight/game.SquareSize-1)*game.SquareSize + int(game.Offset.Y/2)
				food.Position = rl.Vector2{X: float32(rand.Intn(maxX)), Y: float32(rand.Intn(maxY))}
			}

			if !equ {
				break
			}
		}
	}
}
