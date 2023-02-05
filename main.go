package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	screenSize := int32(600)
	rl.InitWindow(screenSize, screenSize, "Space Invaders")

	backgroundTexture := rl.LoadTexture("assets/space_background.png")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTexture(backgroundTexture, 0, 0, rl.White)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
