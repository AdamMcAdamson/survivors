package game

import (
	s "github.com/AdamMcAdamson/survivors/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateWorldCamera() {
	if !rl.CheckCollisionPointRec(rl.GetWorldToScreen2D(Player.Pos, s.World.Camera), s.World.PlayerViewBoundingBox) {
		offset := rl.Vector2Subtract(rl.GetWorldToScreen2D(Player.Pos, s.World.Camera), rl.Vector2{X: s.World.PlayerViewBoundingBox.X, Y: s.World.PlayerViewBoundingBox.Y})
		if offset.X < 0 {
			s.World.Camera.Offset.X -= offset.X
		} else if offset.X > s.World.PlayerViewBoundingBox.Width {
			s.World.Camera.Offset.X -= (offset.X - s.World.PlayerViewBoundingBox.Width)
		}
		if offset.Y < 0 {
			s.World.Camera.Offset.Y -= offset.Y
		} else if offset.Y > s.World.PlayerViewBoundingBox.Height {
			s.World.Camera.Offset.Y -= (offset.Y - s.World.PlayerViewBoundingBox.Height)
		}
	}
}
