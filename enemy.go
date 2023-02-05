package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Enemy struct {
	posX int32
	posY int32
	img  rl.Texture2D
}

// Draws enemies on the screen
func drawEnemies(e []Enemy) []Enemy {
	// uses a temporary array slice to be able to loop over the list while also
	// deleting the enemies that go off screen
	tempSlice := e[:0]
	for i, enemy := range e {
		if enemy.posY <= 0 {
			// deletes by not adding any enemies which have gone off screen to
			// the temporary slice
			continue
		}
		rl.DrawTexture(enemy.img, enemy.posX, enemy.posY, rl.White)
		tempSlice = append(tempSlice, e[i])
	}
	return tempSlice
}
