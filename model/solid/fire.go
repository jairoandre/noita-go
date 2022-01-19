package solid

import (
	"image/color"
	"math/rand"
	"noita-go/model"
	"noita-go/model/gas"
)

type Fire struct {
	LifeSpan float64
	Solid
}

const maxLifespan = 300.0

func NewFire() *Fire {
	fire := Fire{}
	fire.LifeSpan = maxLifespan * (0.85 + rand.Float64()*0.15)
	return &fire
}

func (f *Fire) Update(cell *model.Cell) {
	if f.LifeSpan <= 0 {
		cell.Element = gas.NewSteam()
		return
	}
	f.Solid.Update(cell)
	f.LifeSpan -= 1
}

func (f *Fire) Alpha() color.Alpha {
	alpha := uint8(255.0 * f.LifeSpan / maxLifespan)
	return color.Alpha{A: alpha}
}

func (f *Fire) Color() color.Color {
	rng := rand.Float64()
	if rng < 0.33 {
		return color.White
	} else if rng < 0.66 {
		return color.RGBA{R: 0xff, G: 0xff, A: 0xff}

	} else {
		return color.RGBA{R: 0xff, A: 0xff}
	}
}
