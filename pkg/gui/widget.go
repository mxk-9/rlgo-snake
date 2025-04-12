package gui

import rl "github.com/gen2brain/raylib-go/raylib"

type BaseWidget struct {
	Texture      rl.Texture2D
	SpriteFrames int32
	FrameHeight  float32
	DrawRec      rl.Rectangle

	Pos   rl.Vector2 // x₀; y₀
	Scale rl.Vector2 // dx; dy
	Size  rl.Vector2 // x₁ = (x₀ + scale_x); y₁ = (y₀ + scale_y)

	ShowWidget bool
}

func (b *BaseWidget) GetPos() rl.Vector2 {
	return b.Pos
}

func (b *BaseWidget) SetPos(newPos rl.Vector2) {
	b.Pos = newPos
	b.Scale = rl.NewVector2(
		b.Size.X-b.Pos.X,
		b.Size.Y-b.Pos.Y,
	)
}

func (b *BaseWidget) GetSize() rl.Vector2 {
	return b.Size
}

func (b *BaseWidget) SetSize(newSize rl.Vector2) {
	b.Size = newSize
	b.Scale = rl.NewVector2(
		b.Size.X-b.Pos.X,
		b.Size.Y-b.Pos.Y,
	)
}

func (b *BaseWidget) Unload() {
	rl.UnloadTexture(b.Texture)
}

type Widget interface {
	Draw()
	GetPos() rl.Vector2
	SetPos(newPos rl.Vector2)
	GetSize() rl.Vector2
	SetSize(newSize rl.Vector2)
}
