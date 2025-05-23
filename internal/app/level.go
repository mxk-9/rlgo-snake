package app

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	Borders rl.Rectangle
}

func (l *Level) SetBorders(x0, y0, sizeW, sizeH float32, offset rl.Vector2) {
	l.Borders = rl.Rectangle{
		X:      x0 + offset.X,
		Y:      y0 + offset.Y,
		Width:  sizeW,
		Height: sizeH,
	}
}

func (l *Level) DrawGrid(scr_w, scr_h, square int32, offset rl.Vector2) {
	for i := int32(0); i < scr_w/square+1; i++ {
		rl.DrawLineV(
			rl.Vector2{
				X: float32(square*i) + offset.X/2,
				Y: offset.Y / 2,
			},
			rl.Vector2{
				X: float32(square*i) + offset.X/2,
				Y: float32(scr_h) - offset.Y/2,
			},
			rl.LightGray,
		)
	}

	for i := int32(0); i < scr_h/square+1; i++ {
		rl.DrawLineV(
			rl.Vector2{
				X: offset.X / 2,
				Y: float32(square*i) + offset.Y/2,
			},
			rl.Vector2{
				X: float32(scr_w) - offset.X/2,
				Y: float32(square*i) + offset.Y/2,
			},
			rl.LightGray,
		)
	}
}
