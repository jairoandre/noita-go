package solid

import (
	"image/color"
	"math/rand"
)

type Sand struct {
	Green uint8
	Solid
}

func NewSand() *Sand {
	sand := Sand{}
	sand.Solid.DispersionRate = 3
	sand.Falling = true
	sand.Green = uint8(255 * (rand.Float64()*0.4 + 0.6))
	return &sand
}

func (s *Sand) Color() color.Color {
	return color.RGBA{R: 0xff, G: s.Green, B: 0x00, A: 0xff}
}
