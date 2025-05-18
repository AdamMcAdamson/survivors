package game

import (
	s "github.com/AdamMcAdamson/survivors/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var count = 0

func Update() {
	count++
	conditional(count)
	HandleInput()
	stepGame()
	updateWorldCamera()
	draw()
}

func conditional(count int) {
	if count%100 == 0 {
		var pos = rl.Vector2{X: float32(rl.GetRandomValue(0, 1280)), Y: float32(rl.GetRandomValue(0, 720))}
		var speed = float32(rl.GetRandomValue(100, 500)) / 100
		spawnEnemy(pos, speed)
	}
}

func stepGame() {
	Player.DoPlayerActions()
	updatePositions()
	s.StepCount++
}

func updatePositions() {
	// entities.Player.UpdatePosition()
	for _, e := range entities.Enemies {
		e.moveTowardPlayer()
	}
	for _, p := range entities.Projectiles {
		p.updatePosition()
	}
}
