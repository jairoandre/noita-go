package noita

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"noita-go/utils"
)

type CellType uint8

var CellTypeColors = []color.Color{
	color.Transparent,
	color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff},
	color.RGBA{R: 0x00, G: 0x0d, B: 0xdd, A: 0xff},
	color.RGBA{R: 0x68, G: 0x41, B: 0x32, A: 0xff},
}

const (
	empty  CellType = 0
	sand   CellType = 1
	water  CellType = 2
	ground CellType = 3
)

type Cell struct {
	Tick     int
	Type     CellType
	Pos      utils.Point
	Img      *ebiten.Image
	Alpha    float64
	LifeTime float64
}

func NewCell(x, y float64, img *ebiten.Image, pType CellType) *Cell {
	return &Cell{
		Type:  pType,
		Pos:   utils.Pt(x, y),
		Img:   img,
		Alpha: 1.0,
	}
}

func (p *Cell) Draw(screen *ebiten.Image) {
	var clr = CellTypeColors[p.Type]
	op := &ebiten.DrawImageOptions{}
	r, g, b, a := utils.NormalizeColor(clr)
	op.ColorM.Scale(r, g, b, a*p.Alpha)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(p.Pos.X*scale, p.Pos.Y*scale)
	screen.DrawImage(p.Img, op)
}
