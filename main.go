package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenSize := int32(600)
	playerCoordinates := [2]int32{25, 525}

	rl.InitWindow(screenSize, screenSize, "Space Invaders")

	backgroundImg := rl.LoadTexture("assets/space_background.png")
	playerImg := rl.LoadTexture("assets/player.png")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTexture(backgroundImg, 0, 0, rl.White)
		rl.DrawTexture(playerImg, playerCoordinates[0], playerCoordinates[1], rl.White)

		isInputtingRight := rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight)
		isInputtingLeft := rl.IsKeyDown((rl.KeyA)) || rl.IsKeyDown(rl.KeyLeft)

		if isInputtingRight && !(playerCoordinates[0] > (screenSize - 75)) {
			playerCoordinates[0] += 1
		}
		if isInputtingLeft && !(playerCoordinates[0] <= 25) {
			playerCoordinates[0] -= 1
		}

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
