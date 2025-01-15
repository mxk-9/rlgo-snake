package app

import (
	"rltest/internal/food"
	"rltest/internal/player"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func CollisionWithScreen(head rl.Vector2, screenWidth, screenHeight int32, offset rl.Vector2) (gameOver bool) {
	gameOver = false

	if head.X > float32(screenWidth)-offset.X || head.Y > float32(screenHeight-int32(offset.Y)) ||
		head.X < 0 || head.Y < 0 {
		gameOver = true
		rl.TraceLog(rl.LogInfo, "Snake hits the wall")
	}

	return
}

func SpawnFood(food *food.Food, snake *player.Snake, g *Game) {
	if !food.Active {
		food.Active = true

		food.Position = rl.NewVector2(
			float32(rl.GetRandomValue(0, (g.ScreenWidth/g.SquareSize-1))*g.SquareSize)+g.Offset.X/2,
			float32(rl.GetRandomValue(0, (g.ScreenHeight/g.SquareSize-1))*g.SquareSize)+g.Offset.Y/2,
		)

		for i := int32(0); i < snake.Length; i++ {
			for (food.Position.X == snake.Segments[i].X) && (food.Position.Y == snake.Segments[i].Y) {
				food.Position = rl.NewVector2(
					float32(rl.GetRandomValue(0, (g.ScreenWidth/g.SquareSize-1))*g.SquareSize)+g.Offset.X/2,
					float32(rl.GetRandomValue(0, (g.ScreenHeight/g.SquareSize-1))*g.SquareSize)+g.Offset.Y/2,
				)
				i = 0
			}
		}

	}
}

func CollideYourself(snake *player.Snake) (collide bool) {
	collide = false

	for i := int32(1); i < snake.Length; i++ {
		headX := snake.Segments[0].X
		headY := snake.Segments[0].Y

		if headX == snake.Segments[i].X && headY == snake.Segments[i].Y {
			collide = true
			rl.TraceLog(rl.LogInfo, "Snake eats itself")
			break
		}
	}

	return
}

func SnakeCollidesFood(snake *player.Snake, food *food.Food) (collide bool) {
	collide = false

	snkX := snake.Segments[0].X
	snkY := snake.Segments[0].Y
	fdX := food.Position.X
	fdY := food.Position.Y

	if (snkX < (fdX + float32(food.Size))) &&
		((snkX + snake.Square) > fdX) &&
		(snkY < (fdY + float32(food.Size))) &&
		((snkY + snake.Square) > fdY) {
		snake.Segments[snake.Length] = snake.SegPos[snake.Length-1]
		snake.Length++
		collide = true
	}

	return
}
