package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/martijnspitter/tower-defense/models"
)

type Shooter struct {
	Bullets      []*models.Bullet
	screenWidth  int
	screenHeight int
}

func NewShooter(screenWidth, screenHeigth int) *Shooter {
	return &Shooter{
		screenHeight: screenHeigth,
		screenWidth:  screenWidth,
	}
}

func (s *Shooter) AddBullet(target *models.Enemy) {
	bullet := models.NewBullet(target, s.screenWidth, s.screenHeight)
	s.Bullets = append(s.Bullets, bullet)
}

func (s *Shooter) Update() {
	for i, b := range s.Bullets {
		b.Update()

		// Check if the bullet is outside the screen boundaries
		if b.IsOffScreen(s.screenWidth, s.screenHeight) {
			s.RemoveBullet(i)
		}
	}
}

func (s *Shooter) Draw(screen *ebiten.Image) {
	for _, b := range s.Bullets {
		b.Draw(screen)
	}
}

func (s *Shooter) RemoveBullet(i int) {
	s.Bullets = append(s.Bullets[:i], s.Bullets[i+1:]...)
}
