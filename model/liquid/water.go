package liquid

import "image/color"

type Water struct {
	Liquid
}

func NewWater() Water {
	water := Water{}
	water.DispersionRate = 2
	return water
}

func (w Water) Color() color.Color {
	return color.RGBA{R: 0x20, G: 0x32, B: 0xfe, A: 0xff}
}
