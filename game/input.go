package game

import (
	s "github.com/AdamMcAdamson/survivors/state"
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
	if rl.IsKeyPressed(rl.KeyL) {
		mouseWorldPos := rl.GetScreenToWorld2D(rl.GetMousePosition(), s.World.Camera)
		var angle = -rl.Vector2LineAngle(Player.FaceDir, rl.Vector2{X: 0, Y: 1}) * 360 / rl.Pi
		placeWall(mouseWorldPos.X, mouseWorldPos.Y, 100, 40, angle)
	}
	if rl.IsKeyDown(rl.KeyP) {
		s.World.Camera.Offset.X = 0
		s.World.Camera.Offset.Y = 0
	}

	dir := rl.Vector2Normalize(rl.Vector2{X: dirX, Y: dirY})
	_ = dir
	Player.MoveDir = dir

	if fire {
		Player.FireBullet()
	}

}

func handleMouse() {
	// Player.FaceDir = rl.Vector2Normalize(rl.Vector2Subtract(rl.GetMousePosition(), Player.Pos))
	Player.FaceDir = rl.Vector2Normalize(rl.Vector2Subtract(rl.GetScreenToWorld2D(rl.GetMousePosition(), s.World.Camera), Player.Pos))
	// fmt.Printf("%v || %v || %v\n", s.Player.Pos, s.Player.MoveDir, s.Player.FaceDir)
	// fmt.Println(s.Player.FaceDir)

}
