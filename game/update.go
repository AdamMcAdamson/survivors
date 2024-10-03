package game

func Update() {
	handleInput()
	stepGame()
	draw()
}

func stepGame() {
	updatePositions()
}

func updatePositions() {
	entities.Player.updatePosition()
	for _, p := range entities.Projectiles {
		p.updatePosition()
	}
}
