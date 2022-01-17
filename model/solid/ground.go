package solid

import (
	"image/color"
	"noita-go/model"
)

type Ground struct {
	Solid
}

func NewGround() Ground {
	ground := Ground{}
	ground.Falling = false
	return ground
}

func (g Ground) Update(cell *model.Cell) {
	// do nothing
}

func (g Ground) Color() color.Color {
	return color.RGBA{R: 0x9b, G: 0x76, B: 0x53, A: 0xff}
}
