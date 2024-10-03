package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	entities.Player.draw()
	for _, p := range entities.Projectiles {
		p.draw()
	}

	rl.EndDrawing()
}
