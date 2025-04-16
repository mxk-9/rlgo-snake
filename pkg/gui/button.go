package gui

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	*BaseWidget

	DrawIfPressed bool

	funcOnPress     func()
	funcDrawOnPress func()
}

func NewButton(imagePos rl.Vector2, scale rl.Vector2) (b *Button) {
	b = &Button{}
	b.BaseWidget = NewBaseWidget(imagePos, scale)
	b.DrawIfPressed = false

	return
}

func (b *Button) BindFuncOnPress(funcOnPress func()) {
	b.funcOnPress = funcOnPress
}

func (b *Button) BindDrawFuncOnPress(funcDrawOnPress func()) {
	b.funcDrawOnPress = funcDrawOnPress
}

func (b *Button) Press(mousePos rl.Vector2, mousePressed bool) {
	rect := rl.NewRectangle(b.Pos.X, b.Pos.Y, b.Scale.X, b.Scale.Y)

	if rl.CheckCollisionPointRec(mousePos, rect) && mousePressed {
		if b.funcOnPress != nil {
			b.funcOnPress()
		}
		b.DrawIfPressed = true

		rl.SetTraceLogLevel(rl.LogInfo)
		rl.TraceLog(rl.LogInfo,
			fmt.Sprintf(
				"Button Pressed: %fx%f, draw: %t",
				mousePos.X, mousePos.Y, b.DrawIfPressed,
			),
		)

		if b.funcDrawOnPress != nil {
			b.funcDrawOnPress()
		}
	} else {
		b.DrawIfPressed = false
	}
}

func (b *Button) Draw() {
	if b.ShowWidget {
		rl.DrawTextureRec(b.Texture, b.DrawRec, b.Pos, rl.White)
	}

	if b.funcDrawOnPress != nil {
		b.funcDrawOnPress()
	}
}
