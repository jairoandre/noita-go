package utils

import "math"

type Point struct {
	X    float64
	Y    float64
	XInt int
	YInt int
}

type Vector2D Point

func Vec2(x, y float64) Vector2D {
	return Vector2D{
		X:    x,
		Y:    y,
		XInt: int(x),
		YInt: int(y),
	}
}

func Pt(X, Y float64) Point {
	return Point{
		X: X,
		Y: Y,
	}
}

func (p Point) IntCoords() (int, int) {
	return p.XInt, p.YInt
}

func (p Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func (p Point) Sub(other Point) Point {
	return Point{
		X: p.X - other.X,
		Y: p.Y - other.Y,
	}
}

func (p Point) Mul(k float64) Point {
	return Point{
		X: p.X * k,
		Y: p.Y * k,
	}
}

func (p Point) Length2() float64 {
	return p.X*p.X + p.Y*p.Y
}

func (p Point) Length() float64 {
	return math.Sqrt(p.Length2())
}

func (p Point) SlopeAction(o Point, cb func(x, y float64)) {
	dx := o.X - p.X
	dy := o.Y - p.Y
	absDx := math.Abs(dx)
	absDy := math.Abs(dy)
	if dx == 0 || dy == 0 {
		return
	}
	xDiffIsLarger := absDx > absDy
	stepX := 1.0
	if dx < 0 {
		stepX = -1.0
	}
	stepY := 1.0
	if dy < 0 {
		stepY = -1.0
	}
	longerSideLength := math.Max(absDx, absDy)
	shorterSideLength := math.Min(absDx, absDy)
	slope := shorterSideLength / longerSideLength
	for i := 1.0; i <= longerSideLength; i++ {
		shorterSideIncrease := math.Round(i * slope)
		xIncrease := 0.0
		yIncrease := 0.0
		if xDiffIsLarger {
			xIncrease = i
			yIncrease = shorterSideIncrease
		} else {
			xIncrease = shorterSideIncrease
			yIncrease = i
		}
		toX := p.X + math.Round(xIncrease*stepX)
		toY := p.Y + math.Round(yIncrease*stepY)
		cb(toX, toY)
	}
}
