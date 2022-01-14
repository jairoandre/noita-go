package utils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Particle struct {
	Pos   Point
	Img   *ebiten.Image
	Color color.Color
}

func NewParticle(x, y float64, img *ebiten.Image) *Particle {
	return &Particle{
		Pos: Pt(x, y),
		Img: img,
	}

}
