package scene

import (
	"math"
	"time"

	"github.com/martijnspitter/tower-defense/context"
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
	curTarget    *models.Enemy
	context      *context.Context
}

func NewTowerDefense(screenWidth, screenHeight int, context *context.Context) *TowerDefense {
	tower := models.NewTower(screenWidth, screenHeight)
	shooter := system.NewShooter(screenWidth, screenHeight)

	return &TowerDefense{
		tower:        tower,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		spawntimer:   types.NewTimer(enemySpawnTime),
		shootTimer:   types.NewTimer(time.Duration(shootCooldown)),
		shooter:      shooter,
		context:      context,
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

	// Update current target status
	if td.curTarget != nil {
		targetStillExists := false
		for _, enemy := range td.enemies {
			if enemy.ID == td.curTarget.ID {
				targetStillExists = true
				break
			}
		}
		if !targetStillExists {
			td.curTarget = nil
		}
	}

	td.shootTimer.Update()
	if td.shootTimer.IsReady() {
		td.shootTimer.Reset()

		if td.curTarget == nil && len(td.enemies) > 0 {
			td.curTarget = td.GetClosestEnemy(td.tower.Position())
		}

		if td.curTarget != nil {
			td.shooter.AddBullet(td.curTarget)
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
					if td.curTarget != nil && m.ID.String() == td.curTarget.ID.String() {
						td.curTarget = nil
					}
					td.context.AddPoints(m.Points)
				}
			}
		}
	}

	for i, m := range td.enemies {
		if m.Collider().Intersects(td.tower.Collider()) {
			td.context.RemoveHealth(1)
			td.enemies = append(td.enemies[:i], td.enemies[i+1:]...)
			if td.curTarget != nil && td.curTarget.ID == m.ID {
				td.curTarget = nil
			}
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

func (td *TowerDefense) GetClosestEnemy(towerPos types.Vector) *models.Enemy {
	if len(td.enemies) == 0 {
		return nil
	}

	var closestEnemy *models.Enemy
	minDistance := math.MaxFloat64

	for _, enemy := range td.enemies {
		distance := towerPos.DistanceTo(enemy.Position())
		if distance < minDistance {
			minDistance = distance
			closestEnemy = enemy
		}
	}

	return closestEnemy
}
