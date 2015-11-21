package evogo

import (
	
	"math/rand"
	_"log"

)

var (
	breedMap map[string]BreedFunc = map[string]BreedFunc{
		"twoPointCrossover":twoPointCrossover,
		"cyclicCrossover":cyclicCrossover,
	}
)

func diversityCalc(pop *Population, c *Individual, parents []*Individual) int {
	return pop.DiversityFunc(c, parents)
}

//
// Breed two individuals i and j together, returns a pair of new individual pointers
//
func twoPointCrossover(pop *Population, i, j int, m Mutate) (*Individual,*Individual){
	pa := pop.Individuals[i]
	pb := pop.Individuals[j]

	// Two point crossover (assume lengths are the same for now)
	firstPoint := rand.Intn(len(pa.Chromosome)-1)
	secondPoint := rand.Intn(len(pa.Chromosome)-1)
	if firstPoint > secondPoint {
		firstPoint, secondPoint = secondPoint, firstPoint
	}

	// Make first child
	ca := make([]Gene,0)
	ca = append(ca, pa.Chromosome[:firstPoint]...)
	ca = append(ca, pb.Chromosome[firstPoint:secondPoint]...)
	ca = append(ca, pa.Chromosome[secondPoint:]...)
	// Mutation for first child
	for i, _ := range ca {
		if rand.Float64() < pop.MProb {
			ca[i] = m(ca[i])
		}
	}

	// Make second child
	cb := make([]Gene,0)
	cb = append(cb, pb.Chromosome[:firstPoint]...)
	cb = append(cb, pa.Chromosome[firstPoint:secondPoint]...)
	cb = append(cb, pb.Chromosome[secondPoint:]...)
	for i, _ := range cb {
		if rand.Float64() < pop.MProb {
			ca[i] = m(cb[i])
		}
	}

	//calculate the diversity here (since we still have the parents)
	childA := newIndividualWithGenes(ca)
	childB := newIndividualWithGenes(cb)
	
	if pop.DiversityFunc != nil {
		childA.Diversity = diversityCalc(pop, childA, []*Individual{pa, pb})
		childB.Diversity = diversityCalc(pop, childB, []*Individual{pa, pb})
	}

	return childA,childB 
}

//
// Breed two individuals i and j together using random cycles, returns a pair of new individual pointers
//
func cyclicCrossover(pop *Population, i, j int, m Mutate) (*Individual,*Individual){
	pa := pop.Individuals[i]
	pb := pop.Individuals[j]
	

	if len(pa.Chromosome) != len(pb.Chromosome){
		panic("Cyclic crossover requires same length chromosomes")
	}

	//PARENT A
	// make a random array the length of the chromosome for cyclic crossover
	list := rand.Perm(len(pa.Chromosome))
	cycleArray := make([]int, 0)

	last_num := 0

	for true {
	  val := list[last_num]
	  cycleArray = append(cycleArray, val)
	  if val == 0 {
	  	break
	  }
	  last_num = val
	}

	// Make first child
	ca := make([]Gene,0)
	//now mix the parents based off our cycleArray
	cycle_index := 0
	for i, _ := range pb.Chromosome {
		

		if cycle_index >= len(cycleArray){
			//we ran out of values in our cycle array
			ca = append(ca, pb.Chromosome[i])
			
		} else if i == cycleArray[cycle_index] {
			//grabbing from cycle array
			ca = append(ca, pa.Chromosome[i])
			cycle_index++

		} else {
			//grabbing from other parent array
			ca = append(ca, pb.Chromosome[i])
		}
		
	
	}

	
	//PARENT B
	// make a random array the length of the chromosome for cyclic crossover
	list = rand.Perm(len(pb.Chromosome))
	cycleArray = make([]int, 0)

	last_num = 0

	for true {
	  val := list[last_num]
	  cycleArray = append(cycleArray, val)
	  if val == 0 {
	  	break
	  }
	  last_num = val
	}

	// Make first child
	cb := make([]Gene,0)
	//now mix the parents based off our cycleArray
	cycle_index = 0
	for i, _ := range pb.Chromosome {
		
		if cycle_index >= len(cycleArray){
			//we ran out of values in our cycle array
			cb = append(cb, pa.Chromosome[i])
			
		} else if i == cycleArray[cycle_index] {
			//grabbing from cycle array
			cb = append(cb, pb.Chromosome[i])
			cycle_index++

		} else {
			//grabbing from other parent array
			cb = append(cb, pa.Chromosome[i])
		}
		
	
	}

	return newIndividualWithGenes(ca), newIndividualWithGenes(cb)

}














