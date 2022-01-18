package liquid

import (
	"image/color"
	"math/rand"
)

type Water struct {
	Blue  uint8
	alpha uint8
	Liquid
}

func NewWater() Water {
	water := Water{}
	water.DispersionRate = 3
	water.Blue = uint8(255 * (rand.Float64()*0.2 + 0.8))
	return water
}

func (w Water) Color() color.Color {
	return color.RGBA{R: 0x20, G: 0x32, B: w.Blue, A: 0xff}
}

func (w Water) Alpha() color.Alpha {
	w.alpha = uint8(255 * (rand.Float64()*0.6 + 0.4))
	return color.Alpha{A: w.alpha}
}
