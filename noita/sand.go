package noita

import "math/rand"

func (g *Grid) UpdateSandCell(x, y int, prevCell *Cell) bool {
	cell := g.Get(x, y)
	if cell == nil {
		return false
	}
	if cell.Type == sand || cell.Type == ground {
		return false
	}
	wasWater := cell.Type == water
	// cell.Tick equals to g.Tick means that the particle
	// was updated.
	cell.Tick = g.Tick
	cell.Type = sand
	cell.Alpha = rand.Float64() + 0.2
	if wasWater {
		prevCell.Type = water
		g.UpdateWater(int(prevCell.Pos.X), int(prevCell.Pos.Y))
	}
	return true
}

func (g *Grid) UpdateSand(x, y int) {
	curr := g.Get(x, y)
	if curr == nil || g.Tick == curr.Tick {
		return
	}
	curr.Type = empty
	if g.UpdateSandCell(x, y+1, curr) {
		curr.Tick = g.Tick
		return
	}
	if g.UpdateSandCell(x-1, y+1, curr) {
		curr.Tick = g.Tick
		return
	}
	if g.UpdateSandCell(x+1, y+1, curr) {
		curr.Tick = g.Tick
		return
	}
	curr.Type = sand
}
