package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bullet struct {
	posX      int32
	posY      int32
	velocity  int32
	hitRadius float32
	color     rl.Color
}

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
	currentPosX := int32(75)
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

		tempSlice := bullets[:0]
		for i, bullet := range bullets {
			if bullet.posY <= 0 {
				shouldShoot = true
				continue
			}
			bullets[i].posY = bullets[i].posY - int32(bullet.velocity)
			rl.DrawCircle(bullet.posX, bullet.posY, bullet.hitRadius, bullet.color)
			tempSlice = append(tempSlice, bullets[i])
		}
		bullets = tempSlice

		isInputtingRight := rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight)
		isInputtingLeft := rl.IsKeyDown((rl.KeyA)) || rl.IsKeyDown(rl.KeyLeft)

		if isInputtingRight && !(playerCoordinates[0] > (screenSize - 75)) {
			playerCoordinates[0] += 5
		}
		if isInputtingLeft && !(playerCoordinates[0] <= 25) {
			playerCoordinates[0] -= 5
		}

		if rl.IsKeyDown(rl.KeySpace) && shouldShoot {
			newBullet := Bullet{
				posX:      playerCoordinates[0],
				posY:      playerCoordinates[1],
				velocity:  10,
				hitRadius: float32(10),
				color:     rl.Red,
			}
			bullets = append(bullets, newBullet)
			shouldShoot = false
		}

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
