package gas

import (
	"github.com/mazznoer/colorgrad"
	"image/color"
	"math/rand"
)

type Flame struct {
	MaxLifeSpan float64
	Gas
}

var flameGradient = colorgrad.YlOrRd()

func NewFlame() *Flame {
	const maxLifeSpan = 25
	flame := Flame{}
	flame.LifeSpan = maxLifeSpan * (0.3 + 0.7*rand.Float64())
	flame.MaxLifeSpan = maxLifeSpan
	flame.DispersionRate = 1
	return &flame
}

func (f *Flame) Color() color.Color {
	rate := 1.0 - f.LifeSpan/f.MaxLifeSpan
	clr := flameGradient.At(rate)
	return clr
}

func (f *Flame) Alpha() color.Alpha {
	return color.Alpha{A: 0xff}
}
