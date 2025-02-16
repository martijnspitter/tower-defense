package models

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/martijnspitter/tower-defense/assets"
	"github.com/martijnspitter/tower-defense/types"
)

type Bullet struct {
	target   *Enemy
	position types.Vector
	asset    *ebiten.Image
	movement types.Vector
}

func NewBullet(target *Enemy, screenWidth, screenHeight int) *Bullet {
	asset := assets.MustLoadImage("bullets/bullet.png")

	pos := newVectorForBullet(screenWidth, screenHeight)
	movement := newMovementForBullet(pos, target)

	b := &Bullet{
		target:   target,
		position: pos,
		asset:    asset,
		movement: movement,
	}

	return b
}

func (b *Bullet) Update() {
	b.position.Y += b.movement.Y
	b.position.X += b.movement.X
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	bounds := b.asset.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.2, 0.2)
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(b.position.X, b.position.Y)

	screen.DrawImage(b.asset, op)

}

func (b *Bullet) Collider() types.Rect {
	bounds := b.asset.Bounds()

	return types.NewRect(
		b.position.X,
		b.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}

func (b *Bullet) IsOffScreen(screenWidth, screenHeight int) bool {
	return b.position.X < 0 || b.position.X > float64(screenWidth) || b.position.Y < 0 || b.position.Y > float64(screenHeight)
}

func newVectorForBullet(screenWidth, screenHeight int) types.Vector {
	// Figure out the middle position — the screen center, in this case
	middle := types.Vector{
		X: float64(screenWidth / 2),
		Y: float64(screenHeight / 2),
	}

	// Figure out the spawn position by moving r pixels from the target at the chosen angle
	return types.Vector{
		X: middle.X,
		Y: middle.Y,
	}
}

func newMovementForBullet(pos types.Vector, target *Enemy) types.Vector {
	velocity := float64(3)

	// Direction is the target minus the current position
	direction := types.Vector{
		X: target.position.X - pos.X,
		Y: target.position.Y - pos.Y,
	}

	normalizedDirection := direction.Normalize()

	return types.Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}
}
