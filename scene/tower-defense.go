package scene

import (
	"fmt"
	"time"

	"github.com/martijnspitter/tower-defense/models"
	"github.com/martijnspitter/tower-defense/system"
	"github.com/martijnspitter/tower-defense/types"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	enemySpawnTime = 1 * time.Second
	shootCooldown  = 0.3 * float64(time.Second)
)

type TowerDefense struct {
	tower        *models.Tower
	enemies      []*models.Enemy
	spawntimer   *types.Timer
	screenWidth  int
	screenHeight int
	shooter      *system.Shooter
	shootTimer   *types.Timer
	Health       int
}

func NewTowerDefense(screenWidth, screenHeight int) *TowerDefense {
	tower := models.NewTower(screenWidth, screenHeight)
	shooter := system.NewShooter(screenWidth, screenHeight)

	return &TowerDefense{
		tower:        tower,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		spawntimer:   types.NewTimer(enemySpawnTime),
		shootTimer:   types.NewTimer(time.Duration(shootCooldown)),
		shooter:      shooter,
		Health:       100,
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

	td.shootTimer.Update()
	if td.shootTimer.IsReady() {
		td.shootTimer.Reset()

		if len(td.enemies) > 0 {
			td.shooter.AddBullet(td.enemies[0])
		}
	}

	for _, m := range td.enemies {
		m.Update()
	}

	td.shooter.Update()

	for i, m := range td.enemies {
		for j, b := range td.shooter.Bullets {
			if m.Collider().Intersects(b.Collider()) {
				td.shooter.RemoveBullet(j)
				m.Hit()
				if m.Health == 0 {
					td.enemies = append(td.enemies[:i], td.enemies[i+1:]...)
				}
			}
		}
	}

	for i, m := range td.enemies {
		if m.Collider().Intersects(td.tower.Collider()) {
			td.Health--
			td.enemies = append(td.enemies[:i], td.enemies[i+1:]...)

			fmt.Println("Enemy hit tower")
		}
	}
}

func (td *TowerDefense) Draw(screen *ebiten.Image) {
	td.tower.Draw(screen)

	for _, m := range td.enemies {
		m.Draw(screen)
	}

	td.shooter.Draw(screen)
}

func (td *TowerDefense) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
