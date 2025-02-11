package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/martijnspitter/tower-defense/scene"
)

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Game struct {
	scene        Scene
	screenWidth  int
	screenHeight int
}

func NewGame(screenWidth, screenHeight int) *Game {
	g := &Game{screenWidth: screenWidth, screenHeight: screenHeight}
	g.switchToTD()

	return g
}

func (g *Game) Update() error {
	g.scene.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	if g.screenWidth == 0 || g.screenHeight == 0 {
		return width, height
	}
	return g.screenWidth, g.screenHeight
}

func (g *Game) switchToTD() {
	g.scene = scene.NewTowerDefense()
}
