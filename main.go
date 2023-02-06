package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenSize := int32(600)
	rl.InitWindow(screenSize, screenSize, "Space Invaders")
	gameover := false
	score := int(0)

	backgroundImg := rl.LoadTexture("assets/space_background.png")
	playerImg := rl.LoadTexture("assets/player.png")
	enemyImg := rl.LoadTexture("assets/enemy.png")

	player := createPlayer(playerImg)
	bullets := []Bullet{}
	enemies := []Enemy{}
	enemyCount := 5

	currentPosX := int32(100)
	for enemyCount != 0 {
		newEnemy := Enemy{
			posX:      currentPosX,
			posY:      25,
			img:       enemyImg,
			collision: rl.NewRectangle(float32(currentPosX), 25, 44, 32),
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
		rl.DrawText("Score: "+strconv.Itoa(int(score)), 5, 5, 20, rl.LightGray)
		rl.DrawTexture(player.img, player.posX, player.posY, rl.White)

		if len(enemies) == 0 {
			gameover = true
		}
		if gameover {
			enemies = nil
			bullets = nil
			player.posX = screenSize
			player.posY = screenSize
			rl.UnloadTexture(backgroundImg)
			rl.DrawText("You Won! Your score was: "+strconv.Itoa(int(score)), 125, screenSize/2, 20, rl.LightGray)
		}

		bullets, enemies, score = drawBullets(bullets, enemies, &score)
		enemies = drawEnemies(enemies, player)

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
