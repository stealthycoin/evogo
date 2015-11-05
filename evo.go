package evogo

import (
	"sync"
	"sort"
	"fmt"
	"math/rand"
)

// Fitness evaluation function type
type fitness func(*Individual,[]*Individual) int

// Select two individuals to be mated
// TODO fix last chosen bug
func selectPair(pop* Population) (int,int) {
	// Generate random unique keys then turn into slice
	indicies := make(map[int]bool)
	for len(indicies) < pop.tournament {
		indicies[rand.Intn(len(pop.individuals))] = true
	}
	tourny := make([]int, len(indicies))
	i := 0
	for val, _ := range indicies {
		tourny[i] = val
		i++
	}

	// Sort by index (this also happens to be fitness)
	sort.Ints(tourny)

	// Pick and remove an index from our pool
	pick := func() int {
		// Start at 0 and when the function ends clean out the element that was picked
		i := 0
		defer func() {
			tourny = append(tourny[:i], tourny[i+1:]...)
		}()

		// Loop through and give each individual a tProb chance of being selected
		for _, val := range tourny {
			if rand.Float64() < pop.tProb {
				return val
			}
			i++
		}
		return tourny[len(tourny)-1]
	}

	return pick(), pick()
}

//
// Breed two individuals i and j together based on the population's crossover type
//
func breed(pop *Population, i, j int, m mutate) (*Individual,*Individual) {
	return breedMap[pop.breedMethod](pop, i, j, m)
}


//
// One Generation score all individuals and breed them
//
func breedGeneration(pop *Population, fit fitness, m mutate) *Individual {
	// First score all individuals with the fitness function in seperate goroutines
	var wg sync.WaitGroup
	for i, _ := range pop.individuals {
		wg.Add(1)
		go func (index int) {
			pop.individuals[index].fitness = fit(pop.individuals[index], pop.individuals)
			wg.Done()
		}(i)
	}
	wg.Wait()

	//calculate diversity (optional)
	if pop.diversityFunc != nil {
		//for now the user will iterate through the individuals
		pop.diversityFunc(pop.individuals)	

		// for i, _ := range pop.individuals {
		// 	fmt.Println( pop.individuals[i].Diversity() )
		// }
	}

	// Sort by fitness level
	sort.Sort(pop)

	// Create a pool to hold the next generation
	nextGen := make([]*Individual,0)

	// Copy over elitism winners
	nextGen = append(nextGen, pop.individuals[0:pop.elitism]...)

	// Breed until next gen is full and change population list
	for len(nextGen) < len(pop.individuals) {
		i, j := selectPair(pop)

		// Breed the selected individuals and add the first child to the next gen
		ca, cb := breed(pop, i, j, m)
		nextGen = append(nextGen, ca)

		// Only put in second child if there is room
		if len(nextGen) < len(pop.individuals) {
			nextGen = append(nextGen, cb)
		}

	}
	// Right before losing the last gen save the highest fitness
	fittest := pop.individuals[0]
	pop.individuals = nextGen
	return fittest
}

// Begin the training on a population,
// provide a fitness function and a cutoff
// to stop training when that fitness is reached
func Train(pop *Population, cutoff, maxGen, maxPlateau int, fit fitness, m mutate) *Individual {
	// Set starting genration and best individual
	gen := 0
	var fittest *Individual = nil
	last_fit := cutoff
	plateau := 0

	// Loop until we run out of generataions, or we hit the target, or x generations goes by with no progress
	for gen < maxGen {
		gen++
		fmt.Print("Generation ", gen, "... ")
		fittest = breedGeneration(pop, fit, m)
		fmt.Print("Strongest Candidate: ", fittest.fitness, " ")
		pop.showGenes(fittest)
		fmt.Println()


		// Check to see if we have plateaued
		if fittest.fitness == last_fit {
			plateau++
			if plateau == maxPlateau {
				break
			}
		} else {
			plateau = 0
		}
		last_fit = fittest.fitness

		// Break early if we got a fitness level at or better than our requirement
		if pop.invertFitness && fittest.fitness <= cutoff {
			break
		} else if !pop.invertFitness && fittest.fitness >= cutoff {
			break
		}
	}

	return fittest
}
