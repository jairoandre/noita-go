package model

import (
	"github.com/hajimehoshi/ebiten/v2"
	"noita-go/utils"
)

type Cell struct {
	Tick      int
	Pos       utils.Point
	Scale     float64
	xScaled   float64
	yScaled   float64
	Element   Element
	Img       *ebiten.Image
	Grid      *Grid
	Down      *Cell
	Up        *Cell
	Left      *Cell
	Right     *Cell
	LeftUp    *Cell
	LeftDown  *Cell
	RightUp   *Cell
	RightDown *Cell
}

func NewCell(x, y, scale float64, img *ebiten.Image, grid *Grid) *Cell {
	return &Cell{
		Pos:     utils.Pt(x, y),
		Scale:   scale,
		xScaled: x * scale,
		yScaled: y * scale,
		Grid:    grid,
		Img:     img,
		Element: NewEmpty(),
	}
}

func (p *Cell) SetElement(element Element) {
	p.Element = element
}

func (p *Cell) Update() {
	p.Element.Update(p)
}

func (p *Cell) AlreadyUpdated() bool {
	return p.Tick == p.Grid.Tick
}

func (p *Cell) Draw(screen *ebiten.Image) {
	var clr = p.Element.Color()
	op := &ebiten.DrawImageOptions{}
	r, g, b, a := utils.NormalizeColor(clr)
	op.ColorM.Translate(r, g, b, a)
	op.GeoM.Scale(p.Scale, p.Scale)
	op.GeoM.Translate(p.xScaled, p.yScaled)
	screen.DrawImage(p.Img, op)
}

func (p *Cell) SwapElements(o *Cell) {
	p.Tick = p.Grid.Tick
	o.Tick = p.Grid.Tick
	pElem := p.Element
	p.Element = o.Element
	o.Element = pElem
}
