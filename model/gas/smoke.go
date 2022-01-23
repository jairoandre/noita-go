package gas

import (
	"github.com/mazznoer/colorgrad"
	"image/color"
	"math/rand"
)

type Smoke struct {
	MaxLifeSpan float64
	ColorGrad   colorgrad.Gradient
	Gas
}

var (
	smokeGrad     = colorgrad.Greys()
	smokeLifeSpan = 200.0
)

func NewSmoke() *Smoke {
	smoke := Smoke{}
	smoke.ColorGrad = smokeGrad
	smoke.MaxLifeSpan = smokeLifeSpan
	smoke.LifeSpan = smokeLifeSpan * (0.5 + 0.5*rand.Float64())
	return &smoke
}

func (s *Smoke) Color() color.Color {
	rate := 1.0 - s.LifeSpan/s.MaxLifeSpan
	return s.ColorGrad.At(rate)
}

func (s *Smoke) Alpha() color.Alpha {
	return color.Alpha{A: 0xff}
}
