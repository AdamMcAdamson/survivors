package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var entities Entities
var Player PlayerObj

func Init() {
	Player = PlayerObj{Pos: rl.Vector2{X: 200.0, Y: 200.0}, MoveDir: rl.Vector2{X: 0, Y: 0}, FaceDir: rl.Vector2{X: 0.0, Y: 1.0}, Speed: 10.0}
	loadPlayerTexture()
	entities.Player = &Player
}

func loadPlayerTexture() {
	Player.Texture = rl.LoadTexture("./assets/images/player.png")
}
