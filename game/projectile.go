package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Projectile struct {
	Id    int
	Pos   rl.Vector2
	Dir   rl.Vector2
	Speed float32
}
