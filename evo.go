package evogo

import (
	"sync"
	"sort"
	"fmt"
	"math/rand"
)


// Select two individuals to be mated
// TODO fix last chosen bug
func selectPair(pop* Population) (int,int) {
	// Generate random unique keys then turn into slice
	indicies := make(map[int]bool)
	for len(indicies) < pop.Tournament {
		indicies[rand.Intn(len(pop.Individuals))] = true
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
			if rand.Float64() < pop.TProb {
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
func breed(pop *Population, i, j int, m Mutate) (*Individual,*Individual) {
	return breedMap[pop.BreedMethod](pop, i, j, m)
}

//
//
//
func calcCombinedFitness(pop *Population) {
	
	si := pop.Individuals[0]
	minFitness, maxFitness := si.Fitness, si.Fitness
	minDiversity, maxDiversity := si.Diversity, si.Diversity

	//find min / max diversity and fitness
	for _,in := range pop.Individuals {
		if in.Fitness < minFitness {
			minFitness = in.Fitness
		}
		if in.Fitness > maxFitness {
			maxFitness = in.Fitness
		}
		if in.Diversity < minDiversity {
			minDiversity = in.Diversity
		}
		if in.Diversity > maxDiversity {
			maxDiversity = in.Diversity
		}
	}

	//flip weights if we are inverting fitness
	// if pop.invertFitness {
	// 	temp := minFitness
	// 	minFitness = maxFitness
	// 	maxFitness = temp
	// } 

	for _,in := range pop.Individuals {
		normalizedFitness := float64(in.Fitness - minFitness) / float64(maxFitness - minFitness)
		normalizedDiversity := float64(in.Diversity - minDiversity) / float64(maxDiversity - minDiversity)
		
		// normalizedFitness := float64(in.fitness) / float64(maxFitness)
		// normalizedDiversity := float64(in.diversity) / float64(maxDiversity)

		normalizedFitness *= pop.FitnessWeight
		normalizedDiversity *= pop.DiversityWeight

		if pop.InvertFitness {
			//invert diversity
			normalizedDiversity *= -1
		}	

		in.CombinedFitness = normalizedDiversity + normalizedFitness

	}

	 
}




//
// One Generation score all individuals and breed them
//
func breedGeneration(pop *Population, fit fitness, m Mutate) *Individual {
	// First score all individuals with the fitness function in seperate goroutines
	var wg sync.WaitGroup
	for i, _ := range pop.Individuals {
		wg.Add(1)
		go func (index int) {
			pop.Individuals[index].Fitness = fit(pop.Individuals[index], pop.Individuals)
			wg.Done()
		}(i)
	}
	wg.Wait()

	//calculate combined fitness
	calcCombinedFitness(pop)

	// Sort by fitness level
	sort.Sort(ByFitness(*pop))

	// Create a pool to hold the next generation
	nextGen := make([]*Individual,0)

	// Copy over elitism winners
	nextGen = append(nextGen, pop.Individuals[0:pop.Elitism]...)

	//Create our hobo's and add them
	nextGen = append(nextGen, newIndividuals(pop.Hoboism, pop.MinG, pop.MaxG , pop.NewGene)...)
	
	// Sort by fitness level
	sort.Sort(ByCombinedFitness(*pop))

	// Breed until next gen is full and change population list
	for len(nextGen) < len(pop.Individuals) {
		i, j := selectPair(pop)

		// Breed the selected individuals and add the first child to the next gen
		ca, cb := breed(pop, i, j, m)
		nextGen = append(nextGen, ca)

		// Only put in second child if there is room
		if len(nextGen) < len(pop.Individuals) {
			nextGen = append(nextGen, cb)
		}

	}
	
	// Right before losing the last gen save the highest fitness
	fittest := pop.Individuals[0]
	for _, in := range pop.Individuals {
		if pop.InvertFitness{
			if in.Fitness < fittest.Fitness   {
				fittest = in
			}	
		} else {
			if in.Fitness > fittest.Fitness   {
				fittest = in
			}
		}
	}

	pop.Individuals = nextGen
	return fittest
}

// Begin the training on a population,
// provide a fitness function and a cutoff
// to stop training when that fitness is reached
func Train(pop *Population, cutoff, maxGen, maxPlateau int, fit fitness, m Mutate) *Individual {
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
		fmt.Print("Strongest Candidate -- fitness:", fittest.Fitness, " combinedFitness: ", fittest.CombinedFitness)
		pop.ShowGenes(fittest)
		fmt.Println()

		//run an in loop function
		pop.LoopFunc(pop, gen, plateau)


		// Check to see if we have plateaued
		if (pop.InvertFitness && fittest.Fitness >= last_fit) || (!pop.InvertFitness && fittest.Fitness <= last_fit)  {
			plateau++
			if plateau == maxPlateau {
				break
			}
		} else {
			plateau = 0
		}

		//keep record of last most fit 
		if pop.InvertFitness && fittest.Fitness <= last_fit {
			last_fit = fittest.Fitness
		} else if  !pop.InvertFitness && fittest.Fitness >= last_fit {
			last_fit = fittest.Fitness
		}
		

		// Break early if we got a fitness level at or better than our requirement
		if pop.InvertFitness && fittest.Fitness <= cutoff {
			break
		} else if !pop.InvertFitness && fittest.Fitness >= cutoff {
			break
		}
	}

	return fittest
}
