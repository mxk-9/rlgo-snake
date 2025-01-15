package food

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Food struct {
	Active   bool
	Color    rl.Color
	Position rl.Vector2
	Size     int32
}

func NewFood(size int32, color rl.Color) *Food {
	return &Food{
		Size:   size,
		Color:  color,
		Active: false,
	}
}

func (food *Food) DrawFood() {
	rl.DrawRectangleV(
		food.Position,
		rl.Vector2{X: float32(food.Size), Y: float32(food.Size)},
		food.Color,
	)
}
