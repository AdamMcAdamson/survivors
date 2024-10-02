package game

import (
	s "github.com/AdamMcAdamson/survivors/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Update() {
	handleInput()
	stepGame()
	draw()
}

func stepGame() {
	s.Player.Pos = rl.Vector2Add(s.Player.Pos, rl.Vector2Scale(s.Player.MoveDir, s.Player.Speed))
}
