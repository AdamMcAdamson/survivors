package game

import rl "github.com/gen2brain/raylib-go/raylib"

var count = 0

func Update() {
	count++
	conditional(count)
	handleInput()
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
	updatePositions()
}

func updatePositions() {
	entities.Player.updatePosition()
	for _, e := range entities.Enemies {
		e.moveTowardPlayer()
	}
	for _, p := range entities.Projectiles {
		p.updatePosition()
	}
}
