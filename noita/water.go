package noita

import "math/rand"

func (g *Grid) UpdateWaterCell(x, y int) bool {
	cell := g.Get(x, y)
	if cell != nil && cell.Type == empty {
		cell.Type = water
		cell.Tick = g.Tick
		cell.Alpha = rand.Float64() + 0.2
		return true
	}
	return false
}

func (g *Grid) UpdateWater(x, y int) {
	curr := g.Get(x, y)
	if curr == nil {
		return
	}
	if g.Tick == curr.Tick {
		return
	}
	curr.Tick = g.Tick
	curr.Type = empty
	curr.Type = empty
	if g.UpdateWaterCell(x, y+1) {
		return
	}
	if g.UpdateWaterCell(x-1, y+1) {
		return
	}
	if g.UpdateWaterCell(x+1, y+1) {
		return
	}
	if g.UpdateWaterCell(x-1, y) {
		return
	}
	if g.UpdateWaterCell(x+1, y) {
		return
	}
	curr.Type = water
}
