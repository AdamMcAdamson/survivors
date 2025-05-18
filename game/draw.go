package game

import (
	"fmt"

	s "github.com/AdamMcAdamson/survivors/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func draw() {
	rl.BeginDrawing()
	rl.BeginMode2D(s.World.Camera)

	rl.ClearBackground(rl.RayWhite)

	entities.Player.draw()
	for _, w := range entities.Walls {
		w.draw()
	}

	for _, p := range entities.Projectiles {
		p.draw()
	}
	for _, e := range entities.Enemies {
		e.draw()
	}

	rl.EndMode2D()
	rl.DrawText(fmt.Sprintf("%v %v", s.World.Camera.Offset.X, s.World.Camera.Offset.Y), 10, 10, 24, rl.Gray)
	rl.EndDrawing()
}
