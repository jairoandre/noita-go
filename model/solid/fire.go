package solid

import (
	"github.com/mazznoer/colorgrad"
	"image/color"
	"math/rand"
	"noita-go/model"
	"noita-go/model/gas"
)

type Fire struct {
	MaxLifeSpan float64
	LifeSpan    float64
	Solid
}

var fireGradient = colorgrad.YlOrRd()

const maxFireLifeSpan = 300.0

func NewFire() *Fire {
	fire := Fire{}
	fire.MaxLifeSpan = maxFireLifeSpan
	fire.LifeSpan = maxFireLifeSpan * (0.85 + rand.Float64()*0.15)
	return &fire
}

func (f *Fire) Update(cell *model.Cell) {
	if f.LifeSpan <= 0 {
		cell.Element = model.NewEmpty()
	}
	if cell.Up != nil && cell.Up.Element.Type() == model.EmptyType && rand.Float64() < 0.05 {
		cell.Up.Element = gas.NewFlame()
	}
	f.Solid.Update(cell)
	f.LifeSpan -= 1
}

func (f *Fire) Alpha() color.Alpha {
	return color.Alpha{A: 0xff}
}

func (f *Fire) Color() color.Color {
	rng := rand.Float64()
	idx := 0.0
	if rng < 0.33 {
		idx = 0.45
	} else if rng < 0.66 {
		idx = 0.89
	}
	return fireGradient.At(idx)
}
