package solid

import (
	"image/color"
	"math/rand"
	"noita-go/model"
)

type Ground struct {
	Red uint8
	Solid
}

func NewGround() Ground {
	ground := Ground{}
	ground.Falling = false
	ground.Red = uint8(255 * (rand.Float64()*0.4 + 0.6))
	return ground
}

func (g Ground) Update(cell *model.Cell) {
	// do nothing
}

func (g Ground) Color() color.Color {
	return color.RGBA{R: g.Red, G: 0x76, B: 0x53, A: 0xff}
}
