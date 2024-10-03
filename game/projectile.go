package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Projectile struct {
	Id    int
	Pos   rl.Vector2
	Dir   rl.Vector2
	Speed float32
}

func (p Projectile) draw() {
	var angle = -rl.Vector2LineAngle(p.Dir, rl.Vector2{X: 1, Y: 0}) * 360 / rl.Pi
	rl.DrawRectanglePro(rl.Rectangle{X: p.Pos.X, Y: p.Pos.Y, Width: 10, Height: 2}, rl.Vector2{X: 5, Y: 1}, angle, rl.Green)
}

func (p *Projectile) updatePosition() {
	p.Pos = rl.Vector2Add(p.Pos, rl.Vector2Scale(p.Dir, p.Speed))
}
