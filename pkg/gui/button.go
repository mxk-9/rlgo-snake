package gui

import (
	"embed"
	"fmt"
	"rltest/pkg/assets"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	pos   rl.Vector2
	scale rl.Vector2
	size  rl.Vector2

	texture  rl.Texture2D
	disabled rl.Texture2D
	enabled  rl.Texture2D

	showButton bool

	funcOnPress     func()
	funcDrawOnPress func()
}

func NewButton(imagePos rl.Vector2, scale rl.Vector2) (b *Button) {
	b = &Button{
		pos:   imagePos,
		scale: scale,
		size: rl.Vector2{
			X: imagePos.X + scale.X,
			Y: imagePos.Y + scale.Y,
		},
		showButton: true,
	}

	return
}

func (b *Button) BindFuncOnPress(funcOnPress func()) {
	b.funcOnPress = funcOnPress
}

func (b *Button) BindDrawFuncOnPress(funcDrawOnPress func()) {
	b.funcDrawOnPress = funcDrawOnPress
}

func (b *Button) SetTexture(
	content *embed.FS, imgPath string, rotateDegree int32,
) (err error) {
	var (
		image   *rl.Image
		scale_x int32 = int32(b.scale.X)
		scale_y int32 = int32(b.scale.Y)
	)

	image, err = loadTexture(
		content, imgPath, scale_x, scale_y, rotateDegree,
	)

	if err != nil {
		return
	}

	b.texture = rl.LoadTextureFromImage(image)

	return
}

func (b *Button) Press(mousePos rl.Vector2, mousePressed bool) {
	rect := rl.Rectangle{
		X: b.pos.X, Y: b.pos.Y,
		Width:  float32(b.scale.X),
		Height: float32(b.scale.Y),
	}

	if rl.CheckCollisionPointRec(mousePos, rect) && mousePressed {
		b.funcOnPress()

		rl.SetTraceLogLevel(rl.LogInfo)
		rl.TraceLog(rl.LogInfo,
			fmt.Sprintf("Button Pressed: %fx%f", mousePos.X, mousePos.Y),
		)

	}
}

func (b *Button) Draw() {
	if b.showButton {
		rl.DrawTextureV(b.texture, b.pos, rl.White)
	}
}

func (b *Button) Unload() {
	rl.UnloadTexture(b.disabled)
	rl.UnloadTexture(b.enabled)
}

func (b *Button) Show() {
	b.showButton = true
}

func (b *Button) Hide() {
	b.showButton = false
}

func (b *Button) GetPos() rl.Vector2 {
	return b.pos
}

func (b *Button) SetPos(newPos rl.Vector2) {
	b.pos = newPos
	b.scale = rl.Vector2{
		X: b.size.X - b.pos.X,
		Y: b.size.Y - b.pos.Y,
	}
}

func (b *Button) GetSize() rl.Vector2 {
	return b.size
}

func (b *Button) SetSize(newSize rl.Vector2) {
	b.size = newSize
	b.scale = rl.Vector2{
		X: b.size.X - b.pos.X,
		Y: b.size.Y - b.pos.Y,
	}
}

func loadTexture(content *embed.FS, imgPath string, scale_x, scale_y, rot int32) (
	img *rl.Image, err error,
) {
	if imgPath == "" {
		img = rl.GenImageColor(int(scale_x), int(scale_y), rl.Blank)
		return
	}

	img, err = assets.LoadImage(content, imgPath)
	if err != nil {
		msg := fmt.Sprintf("Failed to load image '%s'", imgPath)
		rl.SetTraceLogLevel(rl.LogError)
		rl.TraceLog(rl.LogError, msg)
		return
	} else {
		rl.SetTraceLogLevel(rl.LogInfo)
		rl.TraceLog(rl.LogInfo, fmt.Sprintf("Whut the FUCK: img is '%v'", img))
		rl.TraceLog(rl.LogInfo, fmt.Sprintf("X: %d\tY: %d", scale_x, scale_y))
	}
	defer rl.UnloadImage(img)

	if scale_x != 0 && scale_y != 0 {
		rl.ImageResizeNN(img, scale_x, scale_y)
	}

	if rot != 0 {
		rl.ImageRotate(img, rot)
	}

	return
}

func (b *Button) PrintDebugInfo() {
	rl.SetTraceLogLevel(rl.LogDebug)

	rl.TraceLog(rl.LogDebug,
		fmt.Sprintf(
			"Button pos : % 4.2fx% 4.2f",
			b.pos.X,
			b.pos.Y,
		),
	)

	rl.TraceLog(rl.LogDebug,
		fmt.Sprintf(
			"Button size: % 4.2fx% 4.2f",
			b.size.X,
			b.size.Y,
		),
	)

	rl.TraceLog(rl.LogDebug,
		fmt.Sprintf(
			"Texture ID : %v", b.texture.ID,
		),
	)

	rl.TraceLog(rl.LogDebug,
		fmt.Sprintf(
			"Enabled ID : %v", b.enabled.ID,
		),
	)

	rl.TraceLog(rl.LogDebug,
		fmt.Sprintf(
			"Disabled ID: %v", b.disabled.ID,
		),
	)
}
