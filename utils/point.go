package utils

import "math"

type Point struct {
	X float64
	Y float64
}

type Vector2D Point

func Vec2(x, y float64) Vector2D {
	return Vector2D{
		X: x,
		Y: y,
	}
}

func Pt(X, Y float64) Point {
	return Point{
		X: X,
		Y: Y,
	}
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
