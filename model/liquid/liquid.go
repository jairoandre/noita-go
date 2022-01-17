package liquid

import (
	"noita-go/model"
	"noita-go/model/solid"
)

type Liquid struct {
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

}

func (l Liquid) LookDown(cell *model.Cell) bool {
	if cell.Down == nil {
		return false
	}
	if cell.Element.Type() != model.EmptyType {
		return false
	}
	return solid.TrySwap(cell, cell.Down)
}
