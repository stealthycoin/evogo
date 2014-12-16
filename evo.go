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
func selectPair(pop* Population) (int,int) {
	// Generate random unique keys then turn into slice
	indicies := make(map[int]bool)
	for len(indicies) < pop.tournament {
		indicies[rand.Intn(len(pop.individuals))] = true
	}
	tourny := make([]int,0)
	for val, _ := range indicies {
		tourny = append(tourny, val)
	}


	// Sort by index (this also happens to be fitness)
	sort.Ints(tourny)

	// Pick first winner
	return_i := -1
	for i, val := range tourny {
		if rand.Float32() < pop.tProb || i == len(tourny)-1{
			return_i = val
			tourny = append(tourny[:i], tourny[i+1:]...)
			break
		}
	}

	// Pick second winner
	return_j := -1
	for i, val := range tourny {
		if rand.Float32() < pop.tProb || i == len(tourny)-1{
			return_j = val
			break
		}
	}

	return return_i, return_j
}

func breed(pop *Population, i, j int, m mutate) (*Individual,*Individual) {
	pa := pop.individuals[i]
	pb := pop.individuals[j]

	// Two point crossover (assume lengths are the same for now)
	firstPoint := rand.Intn(len(pa.chromosome)-1)
	secondPoint := rand.Intn(len(pa.chromosome)-1)
	if firstPoint > secondPoint {
		firstPoint, secondPoint = secondPoint, firstPoint
	}

	// Make first child
	ca := make([]Gene,0)
	ca = append(ca, pa.chromosome[:firstPoint]...)
	ca = append(ca, pb.chromosome[firstPoint:secondPoint]...)
	ca = append(ca, pa.chromosome[secondPoint:]...)

	// Make second child
	cb := make([]Gene,0)
	cb = append(cb, pb.chromosome[:firstPoint]...)
	cb = append(cb, pa.chromosome[firstPoint:secondPoint]...)
	cb = append(cb, pb.chromosome[secondPoint:]...)

	// Mutation for first child
	if rand.Float32() < pop.mProb {
		i := rand.Intn(len(ca))
		ca[i] = m(ca[i])
	}
	// Mutation for second child
	if rand.Float32() < pop.mProb {
		i := rand.Intn(len(cb))
		cb[i] = m(cb[i])
	}

	return newIndividualWithGenes(ca), newIndividualWithGenes(cb)
}

// One Generation score all individuals and breed them
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
func Train(pop *Population, cutoff int, fit fitness, m mutate) {
	maxGenerations := 10000
	gen := 0

	for gen < maxGenerations {
		gen++
		fmt.Print("Generation ", gen, "... ")
		fittest := breedGeneration(pop, fit, m)
		fmt.Print("Strongest Candidate: ", fittest.fitness, " ")
		pop.showGenes(fittest)
		fmt.Println()

		// Break early if we got a fitness level at or better than our requirement
		if pop.invertFitness && fittest.fitness <= cutoff {
			break
		} else if fittest.fitness >= cutoff {
			break
		}
	}
}
