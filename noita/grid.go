package noita

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"noita-go/utils"
	"sync"
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

func (g *Grid) UpdateGoRoutine() {
	g.Tick += 1
	go g.UpdateCells(0, len(g.Cells)/2)
	go g.UpdateCells(len(g.Cells)/2, len(g.Cells))
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

func (g *Grid) DrawRow(screen *ebiten.Image, row []*Cell, total chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, cell := range row {
		if cell.Type == empty {
			continue
		}
		currTotal := <-total
		total <- currTotal + 1
		//cell.Draw(screen)
	}
}

func (g *Grid) Draw(screen *ebiten.Image) {
	total := make(chan int)
	var wg sync.WaitGroup
	for _, row := range g.Cells {
		wg.Add(1)
		go g.DrawRow(screen, row, total, &wg)
	}
	wg.Wait()
	t := <-total
	utils.DebugInfoMessage(screen, fmt.Sprintf("\nTotal particles: %d", t))
}

func (g *Grid) DrawOld(screen *ebiten.Image) {
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
