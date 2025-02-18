package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/martijnspitter/tower-defense/assets"
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
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		if ebiten.TPS() == 0 {
			ebiten.SetTPS(60)
		} else {
			ebiten.TPS(0)
		}
	}
	g.scene.Update()

	if g.context.Health == 0 {
		g.context.ResetHealth()
		g.Reset()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.Draw(screen)
	text.Draw(screen, fmt.Sprintf("%06d", g.context.Points), assets.PointsFont, 50, 50, color.White)
	text.Draw(screen, fmt.Sprintf("%d", g.context.Health), assets.PointsFont, g.screenWidth-150, 50, color.White)
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
