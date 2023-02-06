package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenSize := int32(600)
	rl.InitWindow(screenSize, screenSize, "Space Invaders")
	isVictory := false
	isLoss := false
	score := int(0)

	backgroundImg := rl.LoadTexture("assets/space_background.png")
	playerImg := rl.LoadTexture("assets/player.png")
	enemyImg := rl.LoadTexture("assets/enemy.png")

	player := createPlayer(playerImg)
	enemies := createEnemies(enemyImg)
	bullets := []Bullet{}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTexture(backgroundImg, 0, 0, rl.White)
		rl.DrawText("Score: "+strconv.Itoa(int(score)), 5, 5, 20, rl.LightGray)
		rl.DrawTexture(player.img, player.posX, player.posY, rl.White)

		if len(enemies) == 0 {
			isVictory = true
		}
		if isVictory || isLoss {
			enemies = nil
			bullets = nil
			player.posX = screenSize
			player.posY = screenSize
			if isLoss {
				rl.DrawText("You Lost! Your score was: "+strconv.Itoa(int(score)), 150, screenSize/2, 20, rl.Red)
			} else {
				rl.DrawText("You Won! Your score was: "+strconv.Itoa(int(score)), 150, screenSize/2, 20, rl.Green)
			}
		}

		bullets, enemies, score = drawBullets(bullets, enemies, &score)
		enemies, isLoss = drawEnemies(enemies, player, &isLoss)

		isInputtingRight := rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight)
		if isInputtingRight && !(player.posX > (screenSize - 75)) {
			newPosX := player.posX + 5
			player.posX = newPosX
			player.collision = rl.NewRectangle(float32(newPosX), player.collision.Y, player.collision.Width, player.collision.Height)
		}

		isInputtingLeft := rl.IsKeyDown((rl.KeyA)) || rl.IsKeyDown(rl.KeyLeft)
		if isInputtingLeft && !(player.posX <= 25) {
			newPosX := player.posX - 5
			player.posX = newPosX
			player.collision = rl.NewRectangle(float32(newPosX), player.collision.Y, player.collision.Width, player.collision.Height)
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
