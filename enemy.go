package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Enemy struct {
	posX      int32
	posY      int32
	img       rl.Texture2D
	collision rl.Rectangle
}

// Draws enemies on the screen
func drawEnemies(e []Enemy, p Player) []Enemy {
	// uses a temporary array slice to be able to loop over the list while also
	// deleting the enemies that go off screen
	var tempSlice []Enemy
	for i, enemy := range e {
		if enemy.posY <= 0 {
			// deletes by not adding any enemies which have gone off screen to
			// the temporary slice
			continue
		}
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
	return tempSlice
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
