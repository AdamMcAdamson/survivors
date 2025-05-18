package state

import (
	c "github.com/AdamMcAdamson/survivors/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var World c.WorldObj = c.WorldObj{

	PlayerViewBoundingBox: getBoundingBox(),
	Camera:                rl.NewCamera2D(rl.Vector2{X: 0, Y: 0}, rl.Vector2{X: 0, Y: 0}, 0, 1),
}

func getBoundingBox() rl.Rectangle {
	widthOffset := float32(ScreenWidth) * 0.05
	heightOffset := float32(ScreenWidth) * 0.05

	return rl.NewRectangle(
		widthOffset,
		heightOffset,
		float32(ScreenWidth)-(widthOffset*2),
		float32(ScreenHeight)-(heightOffset*2),
	)
}
