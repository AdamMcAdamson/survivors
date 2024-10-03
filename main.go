package main

import (
	g "github.com/AdamMcAdamson/survivors/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(1280, 720, "raylib [core] example - basic window")

	rl.SetTargetFPS(60)
	g.Init()
	for !rl.WindowShouldClose() {
		g.Update()
	}

	rl.CloseWindow()

}
