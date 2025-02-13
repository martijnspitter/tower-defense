package scene

import (
	"github.com/martijnspitter/tower-defense/models"

	"github.com/hajimehoshi/ebiten/v2"
)

type TowerDefense struct {
	tower *models.Tower
}

func NewTowerDefense(screenWidth, screenHeight int) *TowerDefense {
	tower := models.NewTower(screenWidth, screenHeight)
	return &TowerDefense{
		tower: tower,
	}
}

func (td *TowerDefense) Update() {
	td.tower.Update()
}

func (td *TowerDefense) Draw(screen *ebiten.Image) {
	td.tower.Draw(screen)
}

func (td *TowerDefense) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
