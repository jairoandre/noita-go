package liquid

import (
	"image/color"
	"math/rand"
)

type Water struct {
	Blue uint8
	Liquid
}

func NewWater() Water {
	water := Water{}
	water.DispersionRate = 4
	water.Blue = uint8(255 * (rand.Float64()*0.2 + 0.8))
	return water
}

func (w Water) Color() color.Color {
	return color.RGBA{R: 0x20, G: 0x32, B: w.Blue, A: 0xff}
}
