package main

import (
	"evogo/evo"
	"log"
	"math/rand"
	"time"
)

var (
	letters = " !@#$%^&*(),.[];'.abcdefghijklmnopqrstuvwzyzABCDEFGHIJKLMNOPQRSTUVWZYZ"
	goal = "Hello, World!"
)

// Define a gene in the context of this program
type Gene struct {
	value uint8
}

// Define a function for creating a new random gene
func CreateGene() evo.Gene {
	return Gene{
		value: letters[rand.Intn(len(letters))],
	}
}

// Define a
func fitness(i *evo.Individual) int {
	var f uint8 = 0 // Highest score it can get is 0
	for i, g := range i.Genes() {
		gene := g.(Gene) // Need to typecast it from the eve.Gene interface to our custom Gene type
		f -= (goal[i] - gene.value) // Subtract off how far it is from the target
	}
	return int(f)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	pop := evo.NewPopulation(100, 13, CreateGene)

	log.Println(pop)
}
