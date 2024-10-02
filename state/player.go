package state

import (
	c "github.com/AdamMcAdamson/survivors/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var Player = c.PlayerObj{Pos: rl.Vector2{X: 200.0, Y: 200.0}, MoveDir: rl.Vector2{X: 0, Y: 0}, FaceDir: rl.Vector2{X: 0.0, Y: 1.0}, Speed: 10.0}
var PlayerTexture rl.Texture2D

func Init() {
	PlayerTexture = rl.LoadTexture("./assets/images/player.png")
}
