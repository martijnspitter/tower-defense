package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TowerDefense struct{}

func NewTowerDefense() *TowerDefense {
	return &TowerDefense{}
}

func (td *TowerDefense) Update() {
}

func (td *TowerDefense) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (td *TowerDefense) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
