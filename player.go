package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	posX                int32
	posY                int32
	canShoot            bool
	framesUntilCanShoot int32
	collision           rl.Rectangle
	img                 rl.Texture2D
}

// Creates a player at the default starting coordinates.
func createPlayer(img rl.Texture2D) Player {
	return Player{
		posX:                25,
		posY:                525,
		canShoot:            true,
		framesUntilCanShoot: 0,
		collision:           rl.NewRectangle(25, 525, 50, 48),
		img:                 img,
	}
}
