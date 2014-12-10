package main

import (
	"evogo/evo"
	"fmt"
	"math/rand"
	"time"
	"strings"
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
		value: letters[rand.Intn(len(letters))], // Random rune form our letters allowed
	}
}

func MutateGene(g evo.Gene) evo.Gene {
	gene := g.(Gene)
	i := strings.IndexRune(letters, rune(gene.value))
	i += rand.Intn(5)
	i %= len(letters)
	gene.value = letters[i]
	return gene
}

// Function to display an individual
func ShowGenes(i *evo.Individual) {
	for _, g := range i.Genes() {
		gene := g.(Gene)
		fmt.Printf("%c", rune(gene.value))
	}
}

// Define a
func fitness(i *evo.Individual) int {
	var f int = 0 // Highest score it can get is 0
	for i, g := range i.Genes() {
		gene := g.(Gene) // Need to typecast it from the eve.Gene interface to our custom Gene type
		difference := int(goal[i]) - int(gene.value)
		if difference > 0 {
			difference = -difference
		}
		f += difference  // Subtract off how far it is from the target
	}
	return f
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	// Generate a new populaation with
	// 100 Individuals
	// 13 Genes per individual
	// And our custom CreateGene function used to make the genes
	pop := evo.NewPopulation(100, 13, CreateGene)
	pop.SetShowIndividual(ShowGenes)

	// Call the train function and provide it with
	// The population to train
	// The target fitness at which to stop
	// fitness function to evaluate the fitness of a particular individual
	evo.Train(pop, 0, fitness, MutateGene)
}
