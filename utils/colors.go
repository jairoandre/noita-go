package utils

import "image/color"

func NormalizeColor(clr color.Color) (float64, float64, float64, float64) {
	r, g, b, a := clr.RGBA()
	return float64(r) / 0x10000, float64(g) / 0x10000, float64(b) / 0x10000, float64(a) / 0x10000
}
