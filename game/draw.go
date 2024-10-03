package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	Player.draw()

	rl.EndDrawing()
}
