package snake

import (
	"fmt"
	rl "rltest/pkg/raylib"
)

const InitSnakeLength int = 2

const (
	Red = iota + 1
	Blue
	Green
)

type Snake struct {
	Size     int
	Length   int
	Speed    rl.Vector2
	Segments []rl.Vector2
	SnakePos []rl.Vector2
	Body     rl.Color
	Spots    rl.Color
}

func NewSnake(snakeSkinType int, squareSize int, maxSnakeLength int, offset rl.Vector2) (snake *Snake) {
	snake = &Snake{
		Speed:    rl.Vector2{X: float32(squareSize), Y: 0},
		Size:     squareSize,
		Length:   InitSnakeLength,
		Segments: make([]rl.Vector2, maxSnakeLength),
		SnakePos: make([]rl.Vector2, maxSnakeLength),
	}

	for i := 0; i < maxSnakeLength; i++ {
		snake.Segments[i] = rl.Vector2{
			X: offset.X / 2.0,
			Y: offset.Y / 2.0,
		}
	}

	for i := 0; i < snake.Length; i++ {
		snake.SnakePos[i] = rl.Vector2Zero()
	}

	return
}

func (snake *Snake) SetSkin(snakeSkinType int) {
	switch snakeSkinType {
	case Red:
		snake.Body = rl.PINK
		snake.Spots = rl.RED
	case Blue:
		snake.Body = rl.SKYBLUE
		snake.Spots = rl.BLUE
	case Green:
		snake.Body = rl.GREEN
		snake.Spots = rl.LIME
	default:
		snake.Body = rl.DARKGREEN
		snake.Spots = rl.GREEN
	}
}

func (snake *Snake) RotateSnake() {
	snSize := float32(snake.Size)

	if (rl.IsKeyPressed(rune(rl.KEY_UP)) ||
		rl.IsKeyPressed('W')) && (snake.Speed.Y == 0) {
		snake.Speed = rl.Vector2{X: 0, Y: -snSize}
	}

	if (rl.IsKeyPressed(rune(rl.KEY_DOWN)) ||
		rl.IsKeyPressed('S')) && (snake.Speed.Y == 0) {
		snake.Speed = rl.Vector2{X: 0, Y: snSize}
	}

	if (rl.IsKeyPressed(rune(rl.KEY_RIGHT)) ||
		rl.IsKeyPressed('D')) && (snake.Speed.X == 0) {
		snake.Speed = rl.Vector2{X: snSize, Y: 0}
	}

	if (rl.IsKeyPressed(rune(rl.KEY_LEFT)) ||
		rl.IsKeyPressed('A')) && (snake.Speed.X == 0) {
		snake.Speed = rl.Vector2{X: -snSize, Y: 0}
	}

}

func (snake *Snake) Movement(frame int) {
	for i := 0; i < snake.Length; i++ {
		snake.SnakePos[i] = snake.Segments[i]
	}

	if (frame % 10) == 0 {
		for i := 0; i < snake.Length; i++ {
			if i == 0 {
				snake.Segments[0].X += snake.Speed.X
				snake.Segments[0].Y += snake.Speed.Y
			} else {
				snake.Segments[i] = snake.SnakePos[i-1]
			}
		}
	}
}

func (snake *Snake) DrawSnake() {
	snSize := float32(snake.Size)

	for i := 0; i < snake.Length; i++ {
		rl.DrawRectangleV(snake.Segments[i], rl.Vector2{X: snSize, Y: snSize}, snake.Body)
		rl.DrawRectangleLinesEx(
			rl.Rectangle{
				X:      snake.Segments[i].X,
				Y:      snake.Segments[i].Y,
				Width:  snSize,
				Height: snSize,
			},
			5,
			snake.Spots,
		)
		fmt.Printf("[%2d]%3fx%3f\n", i, snake.Segments[i].X, snake.Segments[i].Y)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()

}

func (snake *Snake) AteItself() (dead bool) {
	dead = false

	for i := 1; i < snake.Length; i++ {
		if (snake.Segments[0].X == snake.Segments[i].X) &&
			(snake.Segments[0].Y == snake.Segments[i].Y) {
			dead = true
			fmt.Println("Snake eats itself")
			break
		}
	}

	return
}
