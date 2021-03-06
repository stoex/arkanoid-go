package components

import (
	"arkanoid/lib/math"

	"github.com/ByteArena/box2d"
)

// Paddle component
type Paddle struct {
	Width  float64
	Height float64
	Body   *box2d.B2Body
}

// Ball component
type Ball struct {
	Radius       float64
	Velocity     float64
	VelocityMult float64 `toml:"velocity_mult"`
	Direction    math.Vector2
	Body         *box2d.B2Body
}

// StickyBall component
type StickyBall struct {
	Period float64
}

// AttractionLine component
type AttractionLine struct{}

// Block component
type Block struct {
	Width  float64
	Height float64
	Health float64
	Body   *box2d.B2Body
}
