package game

import (
	s "github.com/AdamMcAdamson/survivors/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Up int = iota
	Left
	Down
	Right
	Fire
	PlaceWall
	Pickup
	Menu
	Pause

	// Keep last
	UndefinedInput
)

var inputs []bool = make([]bool, UndefinedInput)
var worldMousePos rl.Vector2

func HandleInput() {
	getInputs()
	Player.SetPlayerActions(inputs, worldMousePos)
}

func getInputs() {
	handleMouse()
	handleButtons()
}

func handleButtons() {
	inputs[Up] = rl.IsKeyDown(rl.KeyW)
	inputs[Left] = rl.IsKeyDown(rl.KeyA)
	inputs[Down] = rl.IsKeyDown(rl.KeyS)
	inputs[Right] = rl.IsKeyDown(rl.KeyD)
	inputs[PlaceWall] = rl.IsKeyDown(rl.KeyP)

	inputs[Fire] = rl.IsMouseButtonDown(rl.MouseButtonLeft)
}

func handleMouse() {
	worldMousePos = rl.GetScreenToWorld2D(rl.GetMousePosition(), s.World.Camera)
}
