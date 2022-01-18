package solid

import (
	"image/color"
	"math/rand"
	"noita-go/model"
)

type Solid struct {
	DispersionRate int
	Falling        bool
}

func (s Solid) Weight() int {
	
}

func (s Solid) Type() model.ElementType {
	return model.SolidType
}

func (s Solid) SkipDraw() bool {
	return false
}

func (s Solid) Alpha() color.Alpha {
	return color.Alpha{A: 255}
}

func (s Solid) Update(cell *model.Cell) {
	// When cell ticks is equals to grid tick that means the cell is already updated
	if cell.Tick == cell.Grid.Tick {
		return
	}
	if s.lookDown(cell) {
		return
	}
	s.lookDiagonally(cell)
	// nothing
}

func (s Solid) lookDown(cell *model.Cell) bool {
	return TrySwap(cell, cell.Down)
}

func TrySwap(cell, other *model.Cell) bool {
	if other == nil {
		return false
	}
	if other.Element.Type() == model.SolidType {
		return false
	}
	cell.SwapElements(other)
	return true
}

func (s Solid) lookDiagonally(cell *model.Cell) bool {
	rng := rand.Float64()
	first := cell.LeftDown
	second := cell.RightDown
	if rng < 0.5 {
		first = cell.RightDown
		second = cell.LeftDown
	}
	if TrySwap(cell, first) {
		return true
	}
	if TrySwap(cell, second) {
		return true
	}
	return false
}
