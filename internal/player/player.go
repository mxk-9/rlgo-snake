package player

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Red = iota
	Blue
	Green
)

const InitSnakeLength int32 = 2

type Snake struct {
	Color       rl.Color
	ColorBorder rl.Color
	Speed       rl.Vector2
	Square      float32
	Size        rl.Vector2
	Length      int32
	Segments    []rl.Vector2
	SegPos      []rl.Vector2
	allowMove   bool
}

func NewSnake(square, maxLength, skin int32, offset rl.Vector2) (s *Snake) {
	s = &Snake{
		Square:    float32(square),
		Size:      rl.Vector2{X: float32(square), Y: float32(square)},
		Speed:     rl.Vector2{X: float32(square), Y: 0},
		Length:    InitSnakeLength,
		Segments:  make([]rl.Vector2, maxLength),
		SegPos:    make([]rl.Vector2, maxLength),
		allowMove: true,
	}

	for i := int32(0); i < maxLength; i++ {
		s.Segments[i] = rl.Vector2{
			X: offset.X / 2.0,
			Y: offset.Y / 2.0,
		}
	}

	for i := int32(0); i < s.Length; i++ {
		s.SegPos[i] = rl.Vector2Zero()
	}

	s.setSkin(skin)

	return
}

func (o *Snake) getRect(index int32) (rect rl.Rectangle) {
	rect = rl.Rectangle{
		X:      o.Segments[index].X,
		Y:      o.Segments[index].Y,
		Width:  o.Square,
		Height: o.Square,
	}
	return
}

func (o *Snake) setSkin(skin int32) {
	switch skin {
	case Red:
		o.Color = rl.Pink
		o.ColorBorder = rl.Red
	case Blue:
		o.Color = rl.SkyBlue
		o.ColorBorder = rl.Blue
	case Green:
		o.Color = rl.Green
		o.ColorBorder = rl.Lime
	}
}

func (o *Snake) Draw() {
	for i := int32(0); i < o.Length; i++ {
		rl.DrawRectangleV(o.Segments[i], o.Size, o.Color)
		rl.DrawRectangleLinesEx(o.getRect(i), 5, o.ColorBorder)
	}
}

func (o *Snake) Rotate() {
	if !o.allowMove {
		return
	}

	up := rl.IsKeyPressed(rune(rl.KeyUp)) || rl.IsKeyPressed('W')
	left := rl.IsKeyPressed(rune(rl.KeyLeft)) || rl.IsKeyPressed('A')
	right := rl.IsKeyPressed(rune(rl.KeyRight)) || rl.IsKeyPressed('D')
	down := rl.IsKeyPressed(rune(rl.KeyDown)) || rl.IsKeyPressed('S')

	x := o.Speed.X
	y := o.Speed.Y

	if up && o.Speed.Y <= 0 {
		x = 0
		y = -o.Square
		o.allowMove = false
	} else if left && o.Speed.X <= 0 {
		x = -o.Square
		y = 0
		o.allowMove = false
	} else if right && o.Speed.X >= 0 {
		x = o.Square
		y = 0
		o.allowMove = false
	} else if down && o.Speed.Y >= 0 {
		x = 0
		y = o.Square
		o.allowMove = false
	}

	o.Speed = rl.Vector2{
		X: x, Y: y,
	}

}

func (o *Snake) Move(frame int32, frameTick int32) int32 {
	o.Rotate()

	for i := int32(0); i < o.Length; i++ {
		o.SegPos[i] = o.Segments[i]
	}

	if frame%frameTick == 0 {
		o.allowMove = true
		o.Segments[0].X += o.Speed.X
		o.Segments[0].Y += o.Speed.Y

		for i := int32(1); i < o.Length; i++ {
			o.Segments[i] = o.SegPos[i-1]
		}
	}

	return frame + 1
}
