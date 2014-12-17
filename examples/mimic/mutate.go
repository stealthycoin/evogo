package main

import (
	"math/rand"
	"github.com/stealthycoin/evogo"
)

func MutateGene(g evogo.Gene) evogo.Gene {
	c := g.(Circle)

	// Change the radius
	dr := rand.Intn(7)-3 // [-3, 3]
	c.r += dr
	if c.r < 1 {
		c.r = 1
	}

	// Shift the circle
	dx := rand.Intn(11)-5 // [-5, 5]
	dy := rand.Intn(11)-5 // [-5, 5]
	c.p.X += dx
	c.p.Y += dy


	// Adjust alpha value
	da := rand.Intn(11)-5 // [-5, 5]
	c.a.A += uint8(da)
	if c.a.A < 0 {
		c.a.A = 0
	} else if c.a.A > 255 {
		c.a.A = 255
	}

	// Adjust the color (this is the only way to get colors outside the color profile of the original image)
	dred := uint8(rand.Intn(11)-5) // [-5, 5]
	c.col.R += dred
	if c.col.R < 0 {
		c.col.R = 0
	} else if c.col.A > 255 {
		c.col.R = 255
	}

	dgre := uint8(rand.Intn(11)-5) // [-5, 5]
	c.col.R += dgre
	if c.col.G < 0 {
		c.col.G = 0
	} else if c.col.G > 255 {
		c.col.G = 255
	}

	dblu := uint8(rand.Intn(11)-5) // [-5, 5]
	c.col.R += dblu
	if c.col.B < 0 {
		c.col.B = 0
	} else if c.col.B > 255 {
		c.col.B = 255
	}

	return c
}
