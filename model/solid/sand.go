package solid

import (
	"image/color"
)

type Sand struct {
	Solid
}

func NewSand() Sand {
	sand := Sand{}
	sand.Solid.DispersionRate = 3
	sand.Falling = true
	return sand
}

func (s Sand) Color() color.Color {
	return color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
}
