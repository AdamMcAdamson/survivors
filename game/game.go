package game

type Entities struct {
	NextId      int
	Player      *PlayerObj
	Enemies     []*Enemy
	Projectiles []*Projectile
	Walls       []*Wall
}

func (e *Entities) AddPlayer(p *PlayerObj) {
	p.Id = e.NextId
	e.Player = p
	e.NextId++
}

func (e *Entities) AddProjectile(p *Projectile) {
	p.Id = e.NextId
	e.Projectiles = append(e.Projectiles, p)
	e.NextId++
}

func (e *Entities) AddEnemy(enemy *Enemy) {
	enemy.Id = e.NextId
	e.Enemies = append(e.Enemies, enemy)
	e.NextId++
}
