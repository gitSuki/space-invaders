package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Bullet struct {
	posX      int32
	posY      int32
	velocity  int32
	hitRadius float32
	color     rl.Color
}

// Appends a new bullet to the given array of bullets
func appendBullet(b []Bullet, x int32, y int32) []Bullet {
	newBullet := Bullet{
		posX:      x,
		posY:      y,
		velocity:  10,
		hitRadius: float32(10),
		color:     rl.Red,
	}
	return append(b, newBullet)
}

// Draws the bullets on the screen and calculates their movement.
func drawBullets(b []Bullet) []Bullet {
	// uses a temporary array slice to be able to loop over the list while also
	// deleting the bullets that go off screen
	tempSlice := b[:0]
	for i, bullet := range b {
		if bullet.posY <= 0 {
			// deletes by not adding any bullets which have gone off screen to
			// the temporary slice
			continue
		}
		b[i].posY = b[i].posY - bullet.velocity
		rl.DrawCircle(bullet.posX, bullet.posY, bullet.hitRadius, bullet.color)
		tempSlice = append(tempSlice, b[i])
	}
	return tempSlice
}

// func calculateBulletMovement(bullets []Bullet, i int, velocity int32) {
// 	bullets[i].posY = bullets[i].posY - int32(velocity)
// }
