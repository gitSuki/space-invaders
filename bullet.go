package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bullet struct {
	posX      int32
	posY      int32
	velocity  int32
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
		velocity:  10,
		width:     5,
		height:    15,
		collision: rl.NewRectangle(float32(x+22), float32(y), 5, 15),
		color:     rl.Red,
	}
	return append(b, newBullet)
}

// Draws the bullets on the screen and calculates their movement.
func drawBullets(b []Bullet, e []Enemy) []Bullet {
	// uses a temporary array slice to be able to loop over the list while also
	// deleting the bullets that go off screen
	var tempSlice []Bullet
	for i := range b {
		if b[i].posY <= 0 {
			// deletes by not adding any bullets which have gone off screen to
			// the temporary slice
			continue
		}
		b[i].posY = b[i].posY - b[i].velocity
		b[i].collision = rl.NewRectangle(float32(b[i].posX), float32(b[i].posY), float32(b[i].width), float32(b[i].height))
		rl.DrawRectangle(b[i].posX, b[i].posY, b[i].width, b[i].height, b[i].color)
		tempSlice = append(tempSlice, b[i])
	}
	return tempSlice
}
