package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func detectCollisions(e []Enemy, b []Bullet) ([]Enemy, []Bullet) {
	var newEnemies []Enemy
	var newBullets []Bullet

	for i := range e {
		for j := range b {
			enemyCollision := rl.NewRectangle(float32(e[i].posX), float32(e[i].posY), float32(60), float32(32))
			bulletCollision := rl.NewRectangle(float32(b[j].posX), float32(b[j].posY), float32(b[j].width), float32(b[j].height))
			if rl.CheckCollisionRecs(enemyCollision, bulletCollision) {
				continue
			} else {
				newEnemies = append(newEnemies, e[i])
				newBullets = append(newBullets, b[j])
			}
		}
	}
	return e, b
}
