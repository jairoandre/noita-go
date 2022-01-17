package model

import (
	"image"
	"image/draw"
	"noita-go/utils"
)

type Cell struct {
	Tick      int
	Pos       utils.Point
	Scale     float64
	xScaled   float64
	yScaled   float64
	iXScaled  int
	iYScaled  int
	Element   Element
	Grid      *Grid
	Down      *Cell
	Up        *Cell
	Left      *Cell
	Right     *Cell
	LeftUp    *Cell
	LeftDown  *Cell
	RightUp   *Cell
	RightDown *Cell
	Glowing   bool
	Rect      image.Rectangle
}

func NewCell(x, y, scale float64, grid *Grid) *Cell {
	xScaled := x * scale
	yScaled := y * scale
	iXScaled := int(xScaled)
	iYScaled := int(yScaled)
	iScale := int(scale)
	rect := image.Rect(iXScaled, iYScaled, iXScaled+iScale, iYScaled+iScale)
	return &Cell{
		Pos:      utils.Pt(x, y),
		Scale:    scale,
		xScaled:  x * scale,
		yScaled:  y * scale,
		Grid:     grid,
		Glowing:  true,
		Element:  NewEmpty(),
		iXScaled: iXScaled,
		iYScaled: iYScaled,
		Rect:     rect,
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

func (p *Cell) DrawOnImage(canvas *image.RGBA) {
	uniform := image.NewUniform(p.Element.Color())
	draw.Draw(canvas, p.Rect, uniform, image.Pt(p.iXScaled, p.iYScaled), draw.Src)
}

func (p *Cell) SwapElements(o *Cell) {
	p.Tick = p.Grid.Tick
	o.Tick = p.Grid.Tick
	pElem := p.Element
	p.Element = o.Element
	o.Element = pElem
}
