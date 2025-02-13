package models

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/martijnspitter/tower-defense/assets"
	"github.com/martijnspitter/tower-defense/system"

	"math"
	"math/rand/v2"
)

type Enemy struct {
	position system.Vector
	enemy    *ebiten.Image
	movement system.Vector
}

func NewEnemy(screenWidth, screenHeight int) *Enemy {
	enemies := assets.MustLoadImages("enemies/*.png")
	enemy := enemies[rand.IntN(len(enemies))]

	pos := newVectorForEnemy(screenWidth, screenHeight)
	movement := newMovementForEnemy(screenWidth, screenHeight, pos)

	return &Enemy{
		position: pos,
		enemy:    enemy,
		movement: movement,
	}
}

func (e *Enemy) Update() {
	e.position.X += e.movement.X
	e.position.Y += e.movement.Y
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	bounds := e.enemy.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(e.position.X, e.position.Y)

	screen.DrawImage(e.enemy, op)
}

func newVectorForEnemy(screenWidth, screenHeight int) system.Vector {
	// Figure out the middle position — the screen center, in this case
	middle := system.Vector{
		X: float64(screenWidth / 2),
		Y: float64(screenHeight / 2),
	}

	// The distance from the center the meteor should spawn at — half the width
	r := float64(screenWidth / 2.0)

	// Pick a random angle — 2π is 360° — so this returns 0° to 360°
	angle := rand.Float64() * 2 * math.Pi

	// Figure out the spawn position by moving r pixels from the target at the chosen angle
	return system.Vector{
		X: middle.X + math.Cos(angle)*r,
		Y: middle.Y + math.Sin(angle)*r,
	}
}

func newMovementForEnemy(screenWidth, screenHeight int, pos system.Vector) system.Vector {
	// Randomized	// Figure out the middle position — the screen center, in this case
	middle := system.Vector{
		X: float64(screenWidth / 2),
		Y: float64(screenHeight / 2),
	}

	velocity := 0.25 + rand.Float64()*1.5

	// Direction is the target minus the current position
	direction := system.Vector{
		X: middle.X - pos.X,
		Y: middle.Y - pos.Y,
	}

	// Normalize the vector — get just the direction without the length
	normalizedDirection := direction.Normalize()

	// Multiply the direction by velocity
	return system.Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}
}
