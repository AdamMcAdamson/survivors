package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Enemy struct {
	Id     int
	Health int
	Pos    rl.Vector2
	Speed  float32
}

func (e *Enemy) moveTowardPlayer() {
	e.Pos = rl.Vector2MoveTowards(e.Pos, Player.Pos, e.Speed)
}

func (e Enemy) draw() {
	rl.DrawCircleV(e.Pos, float32(rl.GetRandomValue(10, 20)), rl.Red)
}

func spawnEnemy(position rl.Vector2, speed float32) {
	entities.Enemies = append(entities.Enemies, &Enemy{Id: entities.NextId, Health: 100, Pos: position, Speed: speed})
	entities.NextId++
}
