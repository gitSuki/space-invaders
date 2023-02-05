package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	posX                int32
	posY                int32
	canShoot            bool
	framesUntilCanShoot int32
	img                 rl.Texture2D
}
