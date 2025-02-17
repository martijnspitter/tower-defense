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

	b := &Bullet{
		target:   target,
		position: pos,
		asset:    asset,
	}

	b.movement = b.newMovementForBullet(pos, target)

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
	return types.NewRect(
		b.position.X,
		b.position.Y,
		5,
		5,
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

func (b *Bullet) newMovementForBullet(pos types.Vector, target *Enemy) types.Vector {
	velocity := float64(3)
	bounds := target.asset.Bounds()

	// Enemy center offset (considering 0.5 scale)
	enemyHalfW := float64(bounds.Dx()) * 0.5 / 2
	enemyHalfH := float64(bounds.Dy()) * 0.5 / 2

	// Bullet dimensions (considering 0.2 scale)
	bulletBounds := b.asset.Bounds()
	bulletHalfW := float64(bulletBounds.Dx()) * 0.2 / 2
	bulletHalfH := float64(bulletBounds.Dy()) * 0.2 / 2

	// Direction from bullet center to enemy center
	direction := types.Vector{
		X: (target.Position().X + enemyHalfW) - (pos.X + bulletHalfW),
		Y: (target.Position().Y + enemyHalfH) - (pos.Y + bulletHalfH),
	}

	normalizedDirection := direction.Normalize()

	return types.Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}
}
