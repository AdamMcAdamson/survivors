package config

import rl "github.com/gen2brain/raylib-go/raylib"

type PlayerObj struct {
	Pos     rl.Vector2
	MoveDir rl.Vector2
	FaceDir rl.Vector2
	Speed   float32
}

type WorldObj struct {
	ScreenWidth           int32
	ScreenHeight          int32
	PlayerViewBoundingBox rl.Rectangle
	Camera                rl.Camera2D
}
