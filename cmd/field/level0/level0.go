package level0

import (
	"field/scene"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Level0() (s *scene.Scene) {
	player := rl.Rectangle{
		X:      400,
		Y:      280,
		Width:  40,
		Height: 40,
	}

	camera := rl.Camera2D{
		Target: rl.Vector2{
			X: player.X + 20,
			Y: player.Y + 20,
		},
		Offset: rl.Vector2{
			X: float32(700) / 2,
			Y: float32(700) / 2,
		},
		Rotation: 0,
		Zoom:     1,
	}

	level0 := scene.NewScene()
	level0.Nodes.NewNode("game", nil)
	game := *level0.Nodes.Child["game"]
	game.AddNode("player", player)
	game.AddNode("camera", camera)

	return
}

func Update(s *scene.Scene) {
	cam := s.Nodes.Child["game"].Child["camera"].Object.(*rl.Camera2D)

	if rl.IsKeyDown(rl.KeyA) {
		cam.Rotation--
	} else if rl.IsKeyDown(rl.KeyS) {
		cam.Rotation++
	}

	if cam.Rotation > 40 {
		cam.Rotation = 40
	} else if cam.Rotation < -40 {
		cam.Rotation = -40
	}

	if rl.IsKeyPressed(rl.KeyR) {
		cam.Zoom = 1
		cam.Rotation = 0
	}

	cam.Zoom += float32(rl.GetMouseWheelMove() * 0.05)
}

func Draw(s *scene.Scene) {
	cam := s.Nodes.Child["game"].Child["camera"].Object.(*rl.Camera2D)
	plr := s.Nodes.Child["game"].Child["player"].Object.(*rl.Rectangle)

	rl.BeginMode2D(*cam)

	rl.DrawRectangle(-6000, 320, 13000, 8000, rl.DarkGray)
	rl.DrawRectangleRec(*plr, rl.Red)
	rl.EndMode2D()
}
