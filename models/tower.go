package models

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/martijnspitter/tower-defense/assets"
	"github.com/martijnspitter/tower-defense/system"
)

type Tower struct {
	position system.Vector
	tower    *ebiten.Image
	rotation float64
}

func NewTower(screenWidth, screenHeight int) *Tower {
	tower := assets.MustLoadImage("towers/tower-a.png")

	bounds := tower.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := system.NewVector(float64(screenWidth)/2-halfW, float64(screenHeight)/2-halfH)

	return &Tower{
		position: pos,
		tower:    tower,
	}
}

func (t *Tower) Update() {
	speed := math.Pi / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		t.rotation -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		t.rotation += speed
	}

}

func (p *Tower) Draw(screen *ebiten.Image) {
	bounds := p.tower.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.tower, op)
}
