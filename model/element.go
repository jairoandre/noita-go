package model

import (
	"image/color"
)

type ElementType uint8

const (
	EmptyType  ElementType = 0
	SolidType  ElementType = 1
	LiquidType ElementType = 2
	GasType    ElementType = 3
)

type Element interface {
	Color() color.Color
	Alpha() color.Alpha
	Update(cell *Cell)
	SkipDraw() bool
	Type() ElementType
	Weight() int
}

type Empty struct {
	Element
}

func NewEmpty() Empty {
	return Empty{}
}

func (e Empty) Color() color.Color {
	return color.Transparent
}

func (e Empty) Alpha() color.Alpha {
	return color.Alpha{}
}

func (e Empty) Update(_ *Cell) {
	// nothing
}

func (e Empty) SkipDraw() bool {
	return true
}

func (e Empty) Type() ElementType {
	return EmptyType
}
