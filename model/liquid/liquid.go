package liquid

import (
	"math/rand"
	"noita-go/model"
)

type Liquid struct {
	DispersionRate int
}

func (l Liquid) Weight() int {
	return 1
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
	if l.LookDiagonally(cell) {
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

func (l Liquid) LookDiagonally(cell *model.Cell) bool {
	toLeft := true
	curr := cell.LeftDown
	if rand.Float64() > 0.5 {
		toLeft = false
		curr = cell.RightDown
	}
	if curr == nil || curr.Element.Type() != model.EmptyType {
		if toLeft {
			curr = cell.RightDown
		} else {
			curr = cell.LeftDown
		}
		toLeft = !toLeft
	}
	if curr == nil {
		return false
	}
	for i := 1; i < l.DispersionRate; i++ {
		nextCell := curr.LeftDown
		if !toLeft {
			nextCell = curr.RightDown
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

func (l Liquid) LookSideways(cell *model.Cell) bool {
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
