package gas

import (
	"image/color"
)

type Steam struct {
	Gas
}

const maxLifeSpan = 200
const halfLifeSpan = maxLifeSpan

func NewSteam() *Steam {
	steam := Steam{}
	steam.LifeSpan = maxLifeSpan
	steam.DispersionRate = 1
	return &steam
}

func (s *Steam) Color() color.Color {
	return color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
}

func (s *Steam) Alpha() color.Alpha {
	alpha := uint8(255.0 * float64(s.LifeSpan) / maxLifeSpan)
	return color.Alpha{A: alpha}
}
