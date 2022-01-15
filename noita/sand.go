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
	cell.Type = sand
	cell.Tick = g.Tick
	cell.Alpha = rand.Float64() + 0.2
	if wasWater {
		prevCell.Type = water
		prevCell.Tick = g.Tick - 1
		g.UpdateWater(int(prevCell.Pos.X), int(prevCell.Pos.Y))
	}
	return true
}

func (g *Grid) UpdateSand(x, y int) {
	curr := g.Get(x, y)
	if curr == nil {
		return
	}
	curr.Tick = g.Tick
	curr.Type = empty
	if g.UpdateSandCell(x, y+1, curr) {
		return
	}
	if g.UpdateSandCell(x-1, y+1, curr) {
		return
	}
	if g.UpdateSandCell(x+1, y+1, curr) {
		return
	}
	curr.Type = sand
}