package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerObj struct {
	Id      int
	Pos     rl.Vector2
	MoveDir rl.Vector2
	FaceDir rl.Vector2
	Speed   float32
	Texture rl.Texture2D
}

func (p PlayerObj) draw() {
	var sourceRect = rl.Rectangle{X: 0, Y: 0, Width: 400, Height: 400}
	var destRect = rl.Rectangle{X: p.Pos.X, Y: p.Pos.Y, Width: 80, Height: 80}
	var angle = -rl.Vector2LineAngle(p.FaceDir, rl.Vector2{X: -1, Y: 0}) * 360 / rl.Pi
	rl.DrawTexturePro(p.Texture, sourceRect, destRect, rl.Vector2{X: 25, Y: 40}, angle, rl.White)
}

func (p PlayerObj) FireBullet() {
	entities.Projectiles = append(entities.Projectiles, &Projectile{Id: entities.NextId, Pos: p.Pos, Dir: p.FaceDir, Speed: 10})
	entities.NextId++
}

func (p *PlayerObj) updatePosition() {
	p.Pos = rl.Vector2Add(p.Pos, rl.Vector2Scale(p.MoveDir, p.Speed))
}
