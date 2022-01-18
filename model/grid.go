package model

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"noita-go/utils"
	"sync"
	"sync/atomic"
)

type Grid struct {
	Cells [][]*Cell
	Tick  int
}

func NewGrid() *Grid {
	return &Grid{
		Cells: make([][]*Cell, 0),
		Tick:  0,
	}
}

func (g *Grid) Get(x, y int) *Cell {
	if y < 0 || y >= len(g.Cells) {
		return nil
	}
	row := g.Cells[y]
	if x < 0 || x >= len(row) {
		return nil
	}
	return row[x]
}

func (g *Grid) Update() {
	g.Tick += 1
	startY := len(g.Cells) - 1
	for y := startY; y >= 0; y-- {
		row := g.Cells[y]
		for x := 0; x < len(row); x++ {
			row[x].Update()
		}
	}
}

func (g *Grid) Draw(screen *ebiten.Image, canvas *image.RGBA) {
	var total uint64
	wg := sync.WaitGroup{}
	startY := len(g.Cells) - 1
	for y := startY; y >= 0; y-- {
		row := g.Cells[y]
		wg.Add(1)
		go func() {
			defer wg.Done()
			for x := 0; x < len(row); x++ {
				cell := row[x]
				if cell.Element.SkipDraw() {
					continue
				}
				cell.DrawOnImage(canvas)
				atomic.AddUint64(&total, 1)
			}
		}()
	}
	wg.Wait()
	screen.ReplacePixels(canvas.Pix)
	utils.DebugInfoMessage(screen, fmt.Sprintf("\nTotal particles: %d", total))
}

func (g *Grid) FillCellNeighbor() {
	for y, row := range g.Cells {
		for x, cell := range row {
			cell.Down = g.Get(x, y+1)
			cell.Up = g.Get(x, y-1)
			cell.Left = g.Get(x-1, y)
			cell.Right = g.Get(x+1, y)
			cell.LeftUp = g.Get(x-1, y-1)
			cell.LeftDown = g.Get(x-1, y+1)
			cell.RightUp = g.Get(x+1, y-1)
			cell.RightDown = g.Get(x+1, y+1)
		}
	}
}
