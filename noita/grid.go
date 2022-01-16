package noita

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"noita-go/utils"
)

type Grid struct {
	Cells [][]*Cell
	Tick  int
}

func NewGrid(w, h int) *Grid {
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
			particle := row[x]
			if particle.Tick == g.Tick {
				continue
			}
			switch particle.Type {
			case sand:
				g.UpdateSand(x, y)
			case water:
				g.UpdateWater(x, y)
			default:
				// nothing
			}
		}
	}
}

func (g *Grid) UpdateCells(start, end int) {
	for y := start; y < end; y++ {
		row := g.Cells[y]
		for x := 0; x < len(row); x++ {
			particle := row[x]
			if particle.Tick == g.Tick {
				continue
			}
			switch particle.Type {
			case sand:
				g.UpdateSand(x, y)
			case water:
				g.UpdateWater(x, y)
			default:
				// nothing
			}
		}
	}
}

func (g *Grid) Draw(screen *ebiten.Image) {
	total := 0
	for _, row := range g.Cells {
		for _, cell := range row {
			if cell.Type == empty {
				continue
			}
			total++
			cell.Draw(screen)
		}
	}
	utils.DebugInfoMessage(screen, fmt.Sprintf("\nTotal particles: %d", total))
}
