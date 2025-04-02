package main

import (
	"field/level0"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	lvl := level0.Level0()

	for !rl.WindowShouldClose() {
		level0.Update(lvl)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		level0.Draw(lvl)

		rl.DrawText("Bebra", 640, 10, 20, rl.Red)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
