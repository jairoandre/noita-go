package liquid

import (
	"math/rand"
	"noita-go/model"
)

type Liquid struct {
	DispersionRate int
}

func (l Liquid) Type() model.ElementType {
	return model.LiquidType
}

func (l Liquid) SkipDraw() bool {
	return false
}

func (l Liquid) Update(cell *model.Cell) {
	if cell.AlreadyUpdated() {
		return
	}
	if l.LookDown(cell) {
		return
	}
	l.LookSideways(cell)
}

func (l Liquid) LookDown(cell *model.Cell) bool {
	if cell.Down == nil {
		return false
	}
	if cell.Down.Element.Type() != model.EmptyType {
		return false
	}
	cell.SwapElements(cell.Down)
	return true
}

func (l Liquid) LookSideways(cell *model.Cell) bool {
	toLeft := true
	curr := cell.Left
	if rand.Float64() < 0.6 {
		toLeft = false
		curr = cell.Right
	}

	if curr == nil || curr.Element.Type() != model.EmptyType {
		if toLeft {
			curr = cell.Right
		} else {
			curr = cell.Left
		}
		toLeft = !toLeft
	}
	if curr == nil {
		return false
	}
	for i := 1; i < l.DispersionRate; i++ {
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
