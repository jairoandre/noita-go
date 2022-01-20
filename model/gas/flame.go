package gas

import (
	"github.com/mazznoer/colorgrad"
	"image/color"
)

type Flame struct {
	MaxLifeSpan float64
	Gas
}

var flameGradient = colorgrad.YlOrRd()

func NewFlame() *Flame {
	const maxLifeSpan = 50
	flame := Flame{}
	flame.LifeSpan = maxLifeSpan
	flame.MaxLifeSpan = maxLifeSpan
	flame.DispersionRate = 1
	return &flame
}

func (f *Flame) Color() color.Color {
	rate := 1.0 - float64(f.LifeSpan)/f.MaxLifeSpan
	clr := flameGradient.At(rate)
	return clr
}

func (f *Flame) Alpha() color.Alpha {
	return color.Alpha{A: 0xff}
}
