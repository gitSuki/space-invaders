package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Bullet struct {
	posX      int32
	posY      int32
	speed     int32
	width     int32
	height    int32
	collision rl.Rectangle
	color     rl.Color
}

// Appends a new bullet to the given array of bullets
func appendBullet(b []Bullet, x int32, y int32) []Bullet {
	newBullet := Bullet{
		posX:      x + 22, // offset to make bullet origin at the player's center
		posY:      y,
		speed:     10,
		width:     5,
		height:    15,
		collision: rl.NewRectangle(float32(x+22), float32(y), 5, 15),
		color:     rl.Red,
	}
	return append(b, newBullet)
}

// Draws the bullets on the screen, calculates their movement, and checks for collisions with
// enemies. Which removes both the bullet and the enemy.
func drawBullets(b []Bullet, e []Enemy, score *int) ([]Bullet, []Enemy, int) {
	// uses a temporary array slice to be able to loop over the list while also
	// deleting the bullets that go off screen
	var bulletSlice []Bullet
	var enemySlice []Enemy
	var newScore int
	enemyShouldBeRemoved := false
	for i := range b {
		var collidedEnemyIndex int
		collidedWithEnemy := false

		if b[i].posY <= 0 {
			// deletes by not adding any bullets which have gone off screen to
			// the temporary slice
			continue
		}
		// loops through the list of enemies to detect collisions
		for j := range e {
			collidedWithEnemy = checkBulletCollision(b[i], e[j])
			if collidedWithEnemy {
				collidedEnemyIndex = j
				newScore = *score + 50
				break
			}
		}
		if collidedWithEnemy {
			// removes the enemy and the bullet
			enemySlice = append(e[:collidedEnemyIndex], e[collidedEnemyIndex+1:]...)
			enemyShouldBeRemoved = true
		} else {
			// keeps the bullet in the slice if it hasn't collided
			b[i].posY = b[i].posY - b[i].speed
			b[i].collision = rl.NewRectangle(float32(b[i].posX), float32(b[i].posY), float32(b[i].width), float32(b[i].height))
			rl.DrawRectangle(b[i].posX, b[i].posY, b[i].width, b[i].height, b[i].color)
			bulletSlice = append(bulletSlice, b[i])
		}
	}
	// needs to return a boolean in the case of collision with an enemy
	// in order to update the slice of enemies and the score
	if enemyShouldBeRemoved {
		return bulletSlice, enemySlice, newScore
	} else {
		return bulletSlice, e, *score
	}
}

func checkBulletCollision(b Bullet, e Enemy) bool {
	if rl.CheckCollisionRecs(b.collision, e.collision) {
		return true
	} else {
		return false
	}
}
