package state

import (
	c "github.com/AdamMcAdamson/survivors/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var World c.WorldObj = c.WorldObj{
	ScreenHeight:          720,
	ScreenWidth:           1280,
	PlayerViewBoundingBox: rl.NewRectangle(100, 100, 1280-200, 720-200),
	// CameraOrigin: rl.Vector2Zero(),
	Camera: rl.NewCamera2D(rl.Vector2{X: 0, Y: 0}, rl.Vector2{X: 0, Y: 0}, 0, 1),
}
