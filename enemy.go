package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Enemy struct {
	posX      int32
	posY      int32
	img       rl.Texture2D
	collision rl.Rectangle
}

// Creates a slice of 5 enemies at their default starting coordinates
func createEnemies(img rl.Texture2D) []Enemy {
	var enemies []Enemy
	enemyCount := 5
	currentPosX := int32(100)
	for enemyCount != 0 {
		newEnemy := Enemy{
			posX:      currentPosX,
			posY:      25,
			img:       img,
			collision: rl.NewRectangle(float32(currentPosX), 25, 44, 32),
		}
		enemies = append(enemies, newEnemy)
		enemyCount--
		currentPosX += 100
	}
	return enemies
}

// Draws enemies on the screen and checks for collisions with the player (which triggers game over).
func drawEnemies(e []Enemy, p Player, isLoss *bool) ([]Enemy, bool) {
	// uses a temporary array slice to be able to loop over the list while also
	// deleting the enemies that go off screen
	var tempSlice []Enemy
	enemyHasWon := false
	for i, enemy := range e {
		if enemy.posY <= 0 {
			// deletes by not adding any enemies which have gone off screen to
			// the temporary slice
			continue
		}
		// sets the return boolean to true to trigger a game end state
		// if one of the enemies has collided with the player
		collidedWithPlayer := checkPlayerCollision(p, e[i])
		if collidedWithPlayer {
			enemyHasWon = true
			break
		}

		// adjusts the enemies positioning based on their velocity
		enemyPos := rl.Vector2{
			X: float32(enemy.posX),
			Y: float32(enemy.posY),
		}
		velocity := kinematicSeek(enemy, p)
		newLocation := rl.Vector2Add(enemyPos, velocity)
		e[i].posX = int32(newLocation.X)
		e[i].posY = int32(newLocation.Y)
		e[i].collision = rl.NewRectangle(float32(e[i].posX), float32(e[i].posY), float32(44), float32(32))
		rl.DrawTexture(enemy.img, e[i].posX, e[i].posY, rl.White)
		tempSlice = append(tempSlice, e[i])
	}
	// needs to return a boolean in the case of collision with a player
	// in order to trigger a gameover state
	if enemyHasWon {
		return tempSlice, enemyHasWon
	} else {
		return tempSlice, *isLoss
	}
}

func kinematicSeek(e Enemy, t Player) rl.Vector2 {
	maxSpeed := float32(3)
	enemyPos := rl.Vector2{
		X: float32(e.posX),
		Y: float32(e.posY),
	}
	playerPos := rl.Vector2{
		X: float32(t.posX),
		Y: float32(t.posY),
	}
	velocity := rl.Vector2Subtract(playerPos, enemyPos)
	velocity = rl.Vector2Normalize(velocity)
	velocity = rl.Vector2Scale(velocity, maxSpeed)
	return velocity
}

func checkPlayerCollision(p Player, e Enemy) bool {
	if rl.CheckCollisionRecs(p.collision, e.collision) {
		return true
	} else {
		return false
	}
}
