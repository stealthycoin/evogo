package main

import (
	"image"
	"image/color"
)

/** Struct for circle gene **/
type Circle struct {
	p image.Point        // Center Point
	r int                // Radius
	col color.RGBA       // Color
	a color.Alpha        // Alpha value
}

// Implement the Image interface
func (c *Circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *Circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *Circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < (rr)*(rr) {
		return c.a
	}

	return color.Alpha{0}
}
