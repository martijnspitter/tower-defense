package models

import (
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/martijnspitter/tower-defense/assets"
	"github.com/martijnspitter/tower-defense/types"

	"math"
	"math/rand/v2"
)

type Enemy struct {
	position types.Vector
	asset    *ebiten.Image
	movement types.Vector
	Health   int
	ID       uuid.UUID
}

func NewEnemy(screenWidth, screenHeight int) *Enemy {
	enemies := assets.MustLoadImages("enemies/*.png")
	asset := enemies[rand.IntN(len(enemies))]

	pos := newVectorForEnemy(screenWidth, screenHeight)
	movement := newMovementForEnemy(screenWidth, screenHeight, pos)

	return &Enemy{
		position: pos,
		asset:    asset,
		movement: movement,
		Health:   2,
		ID:       uuid.New(),
	}
}

func (e *Enemy) Update() {
	e.position.Y += e.movement.Y
	e.position.X += e.movement.X
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	bounds := e.asset.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(e.position.X, e.position.Y)

	screen.DrawImage(e.asset, op)
}

func (e *Enemy) Collider() types.Rect {
	return types.NewRect(
		e.position.X,
		e.position.Y,
		10,
		10,
	)
}

func (e *Enemy) Hit() {
	e.Health--
}

func (e *Enemy) Position() types.Vector {
	return e.position
}

func newVectorForEnemy(screenWidth, screenHeight int) types.Vector {
	// Figure out the middle position — the screen center, in this case
	middle := types.Vector{
		X: float64(screenWidth / 2),
		Y: float64(screenHeight / 2),
	}

	// The distance from the center the meteor should spawn at — half the width
	r := float64(screenWidth / 2.0)

	// Pick a random angle — 2π is 360° — so this returns 0° to 360°
	angle := rand.Float64() * 2 * math.Pi

	// Figure out the spawn position by moving r pixels from the target at the chosen angle
	return types.Vector{
		X: middle.X + math.Cos(angle)*r,
		Y: middle.Y + math.Sin(angle)*r,
	}
}

func newMovementForEnemy(screenWidth, screenHeight int, pos types.Vector) types.Vector {
	// Randomized	// Figure out the middle position — the screen center, in this case
	middle := types.Vector{
		X: float64(screenWidth / 2),
		Y: float64(screenHeight / 2),
	}

	velocity := 0.25 + rand.Float64()*1.5

	// Direction is the target minus the current position
	direction := types.Vector{
		X: middle.X - pos.X,
		Y: middle.Y - pos.Y,
	}

	// Normalize the vector — get just the direction without the length
	normalizedDirection := direction.Normalize()

	// Multiply the direction by velocity
	return types.Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}
}
