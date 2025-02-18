package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/martijnspitter/tower-defense/game"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

func main() {
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Tower Defense")
	if err := ebiten.RunGame(game.NewGame(ScreenWidth, ScreenHeight)); err != nil {
		log.Fatal(err)
	}
}
