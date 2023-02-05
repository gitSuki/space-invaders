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
	rl.InitWindow(screenSize, screenSize, "Space Invaders")

	backgroundImg := rl.LoadTexture("assets/space_background.png")
	playerImg := rl.LoadTexture("assets/player.png")
	enemyImg := rl.LoadTexture("assets/enemy.png")

	player := Player{
		posX:                25,
		posY:                525,
		canShoot:            true,
		framesUntilCanShoot: 0,
		img:                 playerImg,
	}
	bullets := []Bullet{}
	enemies := []Enemy{}

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
		rl.DrawTexture(player.img, player.posX, player.posY, rl.White)

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
		if isInputtingRight && !(player.posX > (screenSize - 75)) {
			player.posX += 5
		}

		isInputtingLeft := rl.IsKeyDown((rl.KeyA)) || rl.IsKeyDown(rl.KeyLeft)
		if isInputtingLeft && !(player.posX <= 25) {
			player.posX -= 5
		}

		isInputtingShoot := rl.IsKeyDown(rl.KeySpace)
		if isInputtingShoot && player.canShoot {
			bullets = appendBullet(bullets, player.posX, player.posY)
			player.canShoot = false
			player.framesUntilCanShoot = 30
		}

		if player.framesUntilCanShoot > 0 {
			player.framesUntilCanShoot--
			if player.framesUntilCanShoot == 0 {
				player.canShoot = true
			}
		}

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
