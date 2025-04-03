package gui

import rl "github.com/gen2brain/raylib-go/raylib"

func IsMouseLeftORTapPressed() bool {
	return rl.IsMouseButtonPressed(rl.MouseButtonLeft) ||
		rl.IsGestureDetected(rl.GestureTap)
}
