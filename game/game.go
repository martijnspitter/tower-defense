package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/martijnspitter/tower-defense/context"
	"github.com/martijnspitter/tower-defense/scene"
)

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Game struct {
	scene        *scene.TowerDefense
	screenWidth  int
	screenHeight int
	context      *context.Context
}

func NewGame(screenWidth, screenHeight int) *Game {
	context := context.NewContext()
	g := &Game{screenWidth: screenWidth, screenHeight: screenHeight, context: context}
	g.switchToTD()

	return g
}

func (g *Game) Update() error {
	g.scene.Update()

	if g.context.Health == 0 {
		g.context.ResetHealth()
		g.Reset()
	}

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
	g.scene = scene.NewTowerDefense(g.screenWidth, g.screenHeight, g.context)
}

func (g *Game) Reset() {
	g.scene = scene.NewTowerDefense(g.screenWidth, g.screenHeight, g.context)
}
