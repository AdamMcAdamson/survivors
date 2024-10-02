package game

import (
	s "github.com/AdamMcAdamson/survivors/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	var sourceRect = rl.Rectangle{X: 0, Y: 0, Width: 400, Height: 400}
	var destRect = rl.Rectangle{X: s.Player.Pos.X, Y: s.Player.Pos.Y, Width: 80, Height: 80}
	var angle = -rl.Vector2LineAngle(s.Player.FaceDir, rl.Vector2{X: 1, Y: 0}) * 360 / rl.Pi
	var _ = angle
	// fmt.Println(rect)
	// fmt.Println(angle)

	// var offestPosition = rl.Vector2SubtractValue(s.Player.Pos, 40)
	rl.DrawTexturePro(s.PlayerTexture, sourceRect, destRect, rl.Vector2{X: 25, Y: 40}, angle, rl.White)
	// rl.DrawRectanglePro(rect, rl.Vector2{X: 20, Y: 20}, angle, rl.Blue)
	rl.EndDrawing()
}
