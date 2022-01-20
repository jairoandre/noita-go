package gas

import (
	"math/rand"
	"noita-go/model"
)

type Gas struct {
	DispersionRate int
	LifeSpan       int
}

func (g *Gas) Weight() int {
	return -1
}

func (g *Gas) Type() model.ElementType {
	return model.LiquidType
}

func (g *Gas) SkipDraw() bool {
	return false
}

func (g *Gas) Update(cell *model.Cell) {
	if g.LifeSpan <= 0 {
		cell.Element = model.NewEmpty()
		return
	}
	if cell.AlreadyUpdated() {
		return
	}
	g.LifeSpan -= 1
	if rand.Float64() < 0.8 {
		if g.LookUp(cell) {
			return
		}
		if g.LookDiagonally(cell) {
			return
		}
	} else {
		if g.LookDiagonally(cell) {
			return
		}
		if g.LookUp(cell) {
			return
		}
	}
	g.LookSideways(cell)
}

func (g *Gas) LookUp(cell *model.Cell) bool {
	if cell.Up == nil {
		return false
	}
	if cell.Up.Element.Type() != model.EmptyType {
		return false
	}
	cell.SwapElements(cell.Up)
	return true
}

func (g *Gas) LookDiagonally(cell *model.Cell) bool {
	toLeft := true
	curr := cell.LeftUp
	if rand.Float64() > 0.5 {
		toLeft = false
		curr = cell.RightUp
	}
	if curr == nil || curr.Element.Type() != model.EmptyType {
		if toLeft {
			curr = cell.RightUp
		} else {
			curr = cell.LeftUp
		}
		toLeft = !toLeft
	}
	if curr == nil {
		return false
	}
	for i := 1; i < g.DispersionRate; i++ {
		nextCell := curr.LeftUp
		if !toLeft {
			nextCell = curr.RightUp
		}
		if nextCell == nil {
			break
		}
		if nextCell.Element.Type() != model.EmptyType {
			continue
		}
		curr = nextCell
	}
	if curr.Element.Type() != model.EmptyType {
		return false
	}
	cell.SwapElements(curr)
	return true
}

func (g *Gas) LookSideways(cell *model.Cell) bool {
	toLeft := true
	curr := cell.Left
	if rand.Float64() > 0.5 {
		toLeft = false
		curr = cell.Right
	}
	if curr == nil || curr.Element.Type() != model.EmptyType {
		if toLeft {
			curr = cell.Left
		} else {
			curr = cell.Right
		}
		toLeft = !toLeft
	}
	if curr == nil {
		return false
	}
	for i := 1; i < g.DispersionRate; i++ {
		nextCell := curr.Left
		if !toLeft {
			nextCell = curr.Right
		}
		if nextCell == nil {
			break
		}
		if nextCell.Element.Type() != model.EmptyType {
			continue
		}
		curr = nextCell
	}
	if curr.Element.Type() != model.EmptyType {
		return false
	}
	cell.SwapElements(curr)
	return true
}
