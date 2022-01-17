package model

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"noita-go/utils"
	"sync"
	"time"
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
	for y := 0; y < len(g.Cells); y++ {
		row := g.Cells[y]
		for x := 0; x < len(row); x++ {
			row[x].Update()
		}
	}
}

func (g *Grid) DrawRow(canvas *image.RGBA, row []*Cell, wg *sync.WaitGroup, total chan int) {
	defer wg.Done()
	for x := 0; x < len(row); x++ {
		cell := row[x]
		if cell.Element.SkipDraw() {
			continue
		}
		//cell.Draw(screen)
		cell.DrawOnImage(canvas)
		curr := <-total
		total <- curr + 1
	}

}

func (g *Grid) Draw(screen *ebiten.Image, canvas *image.RGBA) {
	totalChan := make(chan int)
	total := 0
	start := time.Now()
	wg := sync.WaitGroup{}
	for y := 0; y < len(g.Cells); y++ {
		wg.Add(1)
		row := g.Cells[y]
		go g.DrawRow(canvas, row, &wg, totalChan)
	}
	wg.Wait()
	total = <-totalChan
	screen.ReplacePixels(canvas.Pix)
	end := time.Now()
	utils.DebugInfoMessage(screen, fmt.Sprintf("\nTotal particles: %d", total))
	utils.DebugInfoMessage(screen, fmt.Sprintf("\n\n\nTime to Draw: %dms", end.Sub(start).Milliseconds()))
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
