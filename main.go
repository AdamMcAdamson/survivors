package main

import (
	g "github.com/AdamMcAdamson/survivors/game"
	s "github.com/AdamMcAdamson/survivors/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(s.ScreenWidth, s.ScreenHeight, "raylib [core] example - basic window")

	rl.SetTargetFPS(s.FPS)
	g.Init()
	for !rl.WindowShouldClose() {
		g.Update()
	}

	rl.CloseWindow()

}
