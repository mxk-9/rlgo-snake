package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  int32 = 700
	screenHeight int32 = 700
)

func init() {
	rl.InitWindow(screenWidth, screenHeight, "[raylib] - field resizing")
	rl.SetTargetFPS(60)
	rl.SetWindowState(rl.FlagWindowResizable)
	rl.SetWindowState(rl.FlagMsaa4xHint)
}
