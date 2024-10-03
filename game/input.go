package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func handleInput() {
	handleMouse()
	handleButtons()
}

func handleButtons() {
	up := rl.IsKeyDown(rl.KeyW)
	left := rl.IsKeyDown(rl.KeyA)
	down := rl.IsKeyDown(rl.KeyS)
	right := rl.IsKeyDown(rl.KeyD)

	fire := rl.IsMouseButtonDown(rl.MouseButtonLeft)

	var dirX float32 = 0.0
	var dirY float32 = 0.0

	if up {
		dirY -= 1
	}
	if left {
		dirX -= 1
	}
	if down {
		dirY += 1
	}
	if right {
		dirX += 1
	}

	dir := rl.Vector2Normalize(rl.Vector2{X: dirX, Y: dirY})
	_ = dir
	Player.MoveDir = dir

	if fire {
		Player.FireBullet()
	}

}

func handleMouse() {
	Player.FaceDir = rl.Vector2Normalize(rl.Vector2Subtract(rl.GetMousePosition(), Player.Pos))
}
