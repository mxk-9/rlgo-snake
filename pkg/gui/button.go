package gui

import (
	"embed"
	"fmt"
	"rltest/pkg/assets"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	Texture     rl.Texture2D
	Pos         rl.Vector2
	funcOnPress func()
}

func NewButton(content *embed.FS, imagePath string, imagePos rl.Vector2, scale float32,
	funcOnPress func(),
) (b *Button, err error) {
	image, err := assets.LoadImage(content, imagePath)
	if err != nil {
		fmt.Println("ERROR:", "Failed to create button")
		return
	}
	defer rl.UnloadImage(image)

	origW := image.Width
	origH := image.Height

	newW := int32(float32(origW) * scale)
	newH := int32(float32(origH) * scale)

	rl.ImageResize(image, newW, newH)

	b = &Button{
		Texture : rl.LoadTextureFromImage(image),
		Pos: imagePos,
		funcOnPress: funcOnPress,
	}
	
	return
}

func (b *Button) Press(mousePos rl.Vector2, mousePressed bool) {
	rect := rl.Rectangle{
		X: b.Pos.X, Y: b.Pos.Y,
		Width: float32(b.Texture.Width),
		Height: float32(b.Texture.Height),
	}

	if rl.CheckCollisionPointRec(mousePos, rect) && mousePressed {
		b.funcOnPress()
		rl.TraceLog(rl.LogInfo, "Button Pressed!!!!")
	}
}

func (b *Button) Draw() {
	rl.DrawTextureV(b.Texture, b.Pos, rl.White)
}

func (b *Button) Unload() {
	rl.UnloadTexture(b.Texture)
}
