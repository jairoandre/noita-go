package solid

import (
	"github.com/mazznoer/colorgrad"
	"image/color"
	"math/rand"
	"noita-go/model"
	"noita-go/model/gas"
)

type Fire struct {
	LifeSpan float64
	Solid
}

var fireGradient = colorgrad.YlOrRd()

const maxLifespan = 300.0

func NewFire() *Fire {
	fire := Fire{}
	fire.LifeSpan = maxLifespan * (0.85 + rand.Float64()*0.15)
	return &fire
}

func (f *Fire) Update(cell *model.Cell) {
	if f.LifeSpan <= 0 {
		cell.Element = gas.NewFlame()
		return
	}
	f.Solid.Update(cell)
	f.LifeSpan -= 1
}

func (f *Fire) Alpha() color.Alpha {
	return color.Alpha{A: 0xff}
}

func (f *Fire) Color() color.Color {
	rng := rand.Float64() * 0.2
	return fireGradient.At(rng)
}
