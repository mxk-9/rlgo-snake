package app

import (
	"rltest/pkg/player"
	// rl "rltest/pkg/raylib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	statePlay uint8 = iota
	statePause
	stateLose
	stateWin
)

type Game struct {
	FrameTick      int
	State          uint8
	SquareSize     int
	MaxSnakeLength int
	Offset         rl.Vector2
}

func NewGame() (game *Game) {
	game = &Game{}
	game.defaultValues()

	return
}

func (g *Game) defaultValues() {
	g.FrameTick = 10
	g.MaxSnakeLength = 15
	g.SquareSize = 31
	g.State = statePlay
}

func (g *Game) RestartGame(snake *player.Snake) {
	g.defaultValues()
}
