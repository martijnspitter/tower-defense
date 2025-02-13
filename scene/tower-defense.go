package scene

import (
	"time"

	"github.com/martijnspitter/tower-defense/models"
	"github.com/martijnspitter/tower-defense/system"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	enemySpawnTime = 1 * time.Second
)

type TowerDefense struct {
	tower        *models.Tower
	enemies      []*models.Enemy
	spawntimer   *system.Timer
	screenWidth  int
	screenHeight int
}

func NewTowerDefense(screenWidth, screenHeight int) *TowerDefense {
	tower := models.NewTower(screenWidth, screenHeight)
	return &TowerDefense{
		tower:        tower,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		spawntimer:   system.NewTimer(enemySpawnTime),
	}
}

func (td *TowerDefense) Update() {
	td.tower.Update()

	td.spawntimer.Update()
	if td.spawntimer.IsReady() {
		td.spawntimer.Reset()

		m := models.NewEnemy(td.screenWidth, td.screenHeight)
		td.enemies = append(td.enemies, m)
	}

	for _, m := range td.enemies {
		m.Update()
	}
}

func (td *TowerDefense) Draw(screen *ebiten.Image) {
	td.tower.Draw(screen)

	for _, m := range td.enemies {
		m.Draw(screen)
	}
}

func (td *TowerDefense) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
