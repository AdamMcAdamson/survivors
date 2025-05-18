package game

import (
	s "github.com/AdamMcAdamson/survivors/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	UpdatePosition int = iota
	UpdateFaceDir
	TryFireBullet
	TryPlaceWall

	// Keep last
	UndefinedAction
)

var actionsToDo []bool = make([]bool, UndefinedAction)
var actionHandlers = [...]func(*PlayerObj){
	(*PlayerObj).UpdatePosition,
	func(p *PlayerObj) { p.UpdateFaceDir(worldMousePos) },
	(*PlayerObj).TryFireBullet,
	func(p *PlayerObj) { p.TryPlaceWall(worldMousePos) },
}

type PlayerObj struct {
	Id           int
	Pos          rl.Vector2
	MoveDir      rl.Vector2
	FaceDir      rl.Vector2
	Speed        float32
	Texture      rl.Texture2D
	lastFireTime int64
	fireRate     int64
}

func (p PlayerObj) draw() {
	var sourceRect = rl.Rectangle{X: 0, Y: 0, Width: 400, Height: 400}
	var destRect = rl.Rectangle{X: p.Pos.X, Y: p.Pos.Y, Width: 80, Height: 80}
	var angle = -rl.Vector2LineAngle(p.FaceDir, rl.Vector2{X: -1, Y: 0}) * 360 / rl.Pi
	rl.DrawTexturePro(p.Texture, sourceRect, destRect, rl.Vector2{X: 25, Y: 40}, angle, rl.White)
}

func (p *PlayerObj) SetPlayerActions(inputs []bool, worldMousePos rl.Vector2) {
	p.updateMovement(inputs[Up], inputs[Down], inputs[Left], inputs[Right])
	actionsToDo[UpdatePosition] = true
	actionsToDo[UpdateFaceDir] = true
	actionsToDo[TryFireBullet] = inputs[Fire]
	actionsToDo[TryPlaceWall] = inputs[PlaceWall]
}

func (p *PlayerObj) updateMovement(up bool, down bool, left bool, right bool) {
	var dirX float32 = 0.0
	var dirY float32 = 0.0

	if up {
		dirY -= 1
	}
	if left {
		dirX -= 1
	}
	if down {
		dirY += 1
	}
	if right {
		dirX += 1
	}

	dir := rl.Vector2Normalize(rl.Vector2{X: dirX, Y: dirY})
	_ = dir
	Player.MoveDir = dir
}

func (p *PlayerObj) DoPlayerActions() {
	for i, action := range actionsToDo {
		if action {
			actionHandlers[i](p)
		}
	}
}

func (p PlayerObj) FireBullet() {
	entities.Projectiles = append(entities.Projectiles, &Projectile{Id: entities.NextId, Pos: p.Pos, Dir: p.FaceDir, Speed: 10})
	entities.NextId++
}

func (p *PlayerObj) TryFireBullet() {
	if (s.StepCount - p.lastFireTime) >= (int64(s.FPS))/p.fireRate {
		p.FireBullet()
		p.lastFireTime = s.StepCount
	}
}

func (p *PlayerObj) TryPlaceWall(worldMousePos rl.Vector2) {
	conditionToPlaceWall := true
	if conditionToPlaceWall {
		p.PlaceWall(worldMousePos)
	}
}

func (p *PlayerObj) PlaceWall(worldMousePos rl.Vector2) {
	angle := -rl.Vector2LineAngle(Player.FaceDir, rl.Vector2{X: 0, Y: 1}) * 360 / rl.Pi
	placeWall(worldMousePos.X, worldMousePos.Y, 100, 40, angle)
}

func (p *PlayerObj) UpdatePosition() {
	p.Pos = rl.Vector2Add(p.Pos, rl.Vector2Scale(p.MoveDir, p.Speed))
}

func (p *PlayerObj) UpdateFaceDir(screenPos rl.Vector2) {
	p.FaceDir = rl.Vector2Normalize(rl.Vector2Subtract(screenPos, p.Pos))
}
