package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Enemy struct {
	posX int32
	posY int32
	img  rl.Texture2D
}

func main() {
	screenSize := int32(600)
	playerCoordinates := [2]int32{25, 525}
	bullets := []Bullet{}
	enemies := []Enemy{}
	shouldShoot := true

	rl.InitWindow(screenSize, screenSize, "Space Invaders")

	backgroundImg := rl.LoadTexture("assets/space_background.png")
	playerImg := rl.LoadTexture("assets/player.png")
	enemyImg := rl.LoadTexture("assets/enemy.png")

	enemyCount := 5
	currentPosX := int32(100)
	for enemyCount != 0 {
		newEnemy := Enemy{
			posX: currentPosX,
			posY: 25,
			img:  enemyImg,
		}
		enemies = append(enemies, newEnemy)
		enemyCount--
		currentPosX += 100
	}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTexture(backgroundImg, 0, 0, rl.White)
		rl.DrawTexture(playerImg, playerCoordinates[0], playerCoordinates[1], rl.White)

		tempSliceE := enemies[:0]
		for i, enemy := range enemies {
			if enemy.posY <= 0 {
				continue
			}
			rl.DrawTexture(enemy.img, enemy.posX, enemy.posY, rl.White)
			tempSliceE = append(tempSliceE, enemies[i])
		}
		enemies = tempSliceE

		bullets = drawBullets(bullets)

		isInputtingRight := rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight)
		if isInputtingRight && !(playerCoordinates[0] > (screenSize - 75)) {
			playerCoordinates[0] += 5
		}

		isInputtingLeft := rl.IsKeyDown((rl.KeyA)) || rl.IsKeyDown(rl.KeyLeft)
		if isInputtingLeft && !(playerCoordinates[0] <= 25) {
			playerCoordinates[0] -= 5
		}

		isInputtingShoot := rl.IsKeyDown(rl.KeySpace)
		if isInputtingShoot && shouldShoot {
			bullets = appendBullet(bullets, playerCoordinates[0], playerCoordinates[1])
		}

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
