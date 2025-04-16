package gui

import (
	"embed"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"rltest/pkg/assets"
)

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

func NewBaseWidget(pos, scale rl.Vector2) (bw *BaseWidget) {
	bw = &BaseWidget{
		Pos:   pos,
		Scale: scale,
		Size: rl.NewVector2(
			pos.X+scale.X, pos.Y+scale.Y,
		),
		ShowWidget: true,
	}

	return
}

func (b *BaseWidget) SetTexture(
	content *embed.FS, imagePath string, frames int32, rotateDegree int32,
) (err error) {
	if frames <= 0 {
		rl.SetTraceLogLevel(rl.LogError)
		rl.TraceLog(rl.LogError,
			fmt.Sprintf("Frame count cannot be less than 1! Got: '%d'", frames),
		)
		err = fmt.Errorf("FramesLessThanOne")
		return
	}

	image, err := assets.LoadImage(content, imagePath)
	if err != nil {
		rl.SetTraceLogLevel(rl.LogError)
		errMsg := fmt.Sprintf("Failed to create button: %v", err)
		rl.TraceLog(rl.LogError, errMsg)
	}
	defer rl.UnloadImage(image)

	b.SpriteFrames = frames

	if b.Scale.X <= 0 || b.Scale.Y <= 0 {
		rl.SetTraceLogLevel(rl.LogError)
		rl.TraceLog(rl.LogError,
			fmt.Sprintf(
				"Scale cannot be less that 0! Got: '%3.2fx%3.2f'",
				b.Scale.X, b.Scale.Y,
			),
		)
	}
	rl.ImageResizeNN(image, int32(b.Scale.X), int32(b.Scale.Y)*frames)

	if rotateDegree != 0 {
		rl.ImageRotate(image, rotateDegree)
	}

	b.Texture = rl.LoadTextureFromImage(image)
	b.FrameHeight = float32(b.Texture.Height) / float32(frames)
	b.DrawRec = rl.NewRectangle(0, 0, float32(b.Texture.Width), b.FrameHeight)

	rl.SetTraceLogLevel(rl.LogInfo)
	rl.TraceLog(
		rl.LogInfo,
		fmt.Sprintf(
			"Widget info: DrawRec: %v; Texture: %v", b.DrawRec, b.Texture,
		),
	)
	return
}

func (b *BaseWidget) SetTextureFrame(frame int32) {
	if frame < 1 || frame > b.SpriteFrames {
		rl.SetTraceLogLevel(rl.LogWarning)
		rl.TraceLog(rl.LogWarning,
			fmt.Sprintf("Frame number is out of bounds. Got: %d, Max: %d",
				frame, b.SpriteFrames,
			))
	} else {
		b.DrawRec = rl.NewRectangle(0, b.FrameHeight*float32(frame-1),
			float32(b.Texture.Width), b.FrameHeight,
		)
	}
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
