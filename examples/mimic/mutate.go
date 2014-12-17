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

	return c
}
