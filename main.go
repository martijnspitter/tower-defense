package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/martijnspitter/tower-defense/game"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Tower Defense")
	if err := ebiten.RunGame(game.NewGame(ScreenWidth, ScreenHeight)); err != nil {
		log.Fatal(err)
	}
}
