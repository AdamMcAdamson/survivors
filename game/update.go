package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Update() {
	handleInput()
	stepGame()
	draw()
}

func stepGame() {
	Player.Pos = rl.Vector2Add(Player.Pos, rl.Vector2Scale(Player.MoveDir, Player.Speed))
}
