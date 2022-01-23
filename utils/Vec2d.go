package utils

import "math/rand"

type Vec2d struct {
	X float64
	Y float64
}

func RndVec2d() Vec2d {
	return Vec2d{
		X: 2 * (rand.Float64() - 0.5),
		Y: 2 * (rand.Float64() - 0.5),
	}
}

func ZeroVec2d() Vec2d {
	return Vec2d{0, 0}
}

func Vec2dXY(X, Y float64) Vec2d {
	return Vec2d{X: X, Y: Y}
}

func (v Vec2d) Add(o Vec2d) Vec2d {
	return Vec2d{
		X: v.X + o.X,
		Y: v.Y + o.Y,
	}
}

func (v Vec2d) Sub(o Vec2d) Vec2d {
	return Vec2d{
		X: v.X - o.X,
		Y: v.Y - o.Y,
	}
}

func (v Vec2d) Mul(f float64) Vec2d {
	return Vec2d{
		X: v.X * f,
		Y: v.Y * f,
	}
}

func (v Vec2d) Div(f float64) Vec2d {
	return Vec2d{
		X: v.X / f,
		Y: v.Y / f,
	}
}

func (v Vec2d) ToPt() Point {
	return Point{
		X: v.X,
		Y: v.Y,
	}
}
