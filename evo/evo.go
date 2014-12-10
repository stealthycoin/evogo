package evo

import (
	"sync"
	"fmt"
)

// Fitness evaluation function type
type fitness func(*Individual) int


// One Generation score all individuals and breed them
func breedGeneration(pop *Population, fit fitness) {
	// First score all individuals with the fitness function in seperate goroutines
	var wg sync.WaitGroup
	for _, individual := range pop.individuals {
		wg.Add(1)
		go func () {
			individual.fitness = fit(individual)
		}()
	}
	wg.Wait()

	// Sort by fitness level higher is better
}

// Begin the training on a population,
// provide a fitness function and a cutoff
// to stop training when that fitness is reached
func Train(pop *Population, cutoff int, fit fitness) {
	maxGenerations := 100
	gen := 0

	for gen < maxGenerations {
		gen++
		fmt.Println("Generation ", gen)
		breedGeneration(pop, fit)

	}
}
